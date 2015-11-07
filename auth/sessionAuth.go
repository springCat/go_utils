package auth

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/context"
	"net/http"
	"encoding/gob"
)

type (
	AuthConf struct {
		Login           AuthHandle
		Logout          AuthHandle
		LoginSuccess    AuthHandle
		LoginFailed     AuthHandle
		UnAuthenticated AuthHandle
		IsAuthenticated AuthHandle
		Session         AuthSession
		Redis           AuthRedis
		LoginUser       User
	}

	AuthHandle struct {
		url    string
		handle gin.HandlerFunc
	}

	AuthSession struct {
		CookieKey  string
		SessionKey string
		Secret     string
		MaxAge     int
	}

	AuthRedis struct {
		Size     int
		Network  string
		Address  string
		Password string
	}
)

var authConf AuthConf

func NewDefaultRedisSession(engine *gin.Engine, user User, MaxAge int, Secret string, redisServerAddress string, redisServerPassword string) {
	conf := AuthConf{
		Login: AuthHandle{
			url:    "/login",
			handle: login,
		},
		Logout: AuthHandle{
			url:    "/logout",
			handle: DefaultLogout,
		},
		LoginSuccess: AuthHandle{
			url:    "/LoginSuccess",
			handle: DefaultLoginSuccess,
		},
		LoginFailed: AuthHandle{
			url:    "/loginFailed",
			handle: DefaultLoginFailed,
		},
		UnAuthenticated: AuthHandle{
			url:    "/unAuthenticated",
			handle: DefaultUnAuthenticate,
		},
		IsAuthenticated: AuthHandle{
			url:    "/isAuthenticated",
			handle: DefaultIsAuthenticated,
		},
		Session: AuthSession{
			CookieKey:  "GSESSIONID",
			SessionKey: "AUTHENTICId",
			Secret:     Secret,
			MaxAge:     MaxAge,
		},
		Redis: AuthRedis{
			Size:     10,
			Network:  "tcp",
			Address:  redisServerAddress,
			Password: redisServerPassword,
		},
		LoginUser: user,
	}
	NewCookieSession(engine, conf)
}

func NewDefaultCookieSession(engine *gin.Engine, user User, MaxAge int, Secret string) {
	conf := AuthConf{
		Login: AuthHandle{
			url:    "/login",
			handle: login,
		},
		Logout: AuthHandle{
			url:    "/logout",
			handle: DefaultLogout,
		},
		LoginSuccess: AuthHandle{
			url:    "/LoginSuccess",
			handle: DefaultLoginSuccess,
		},
		LoginFailed: AuthHandle{
			url:    "/loginFailed",
			handle: DefaultLoginFailed,
		},
		UnAuthenticated: AuthHandle{
			url:    "/unAuthenticated",
			handle: DefaultUnAuthenticate,
		},
		IsAuthenticated: AuthHandle{
			url:    "/isAuthenticated",
			handle: DefaultIsAuthenticated,
		},
		Session: AuthSession{
			CookieKey:  "GSESSIONID",
			SessionKey: "AUTHENTICId",
			Secret:     Secret,
			MaxAge:     MaxAge,
		},
		LoginUser: user,
	}
	NewCookieSession(engine, conf)
}

func NewCookieSession(engine *gin.Engine, conf AuthConf) {
	store := sessions.NewCookieStore([]byte(conf.Session.Secret))
	newSession(engine, conf, store)
}

func NewRedisSession(engine *gin.Engine, conf AuthConf) {
	store, err := sessions.NewRedisStore(conf.Redis.Size, conf.Redis.Network, conf.Redis.Address, conf.Redis.Password, []byte(conf.Session.Secret))
	if err != nil {
		panic(err)
	}
	newSession(engine, conf, store)
}

func newSession(engine *gin.Engine, conf AuthConf, store sessions.CookieStore) {
	engine.Use(sessions.Sessions(conf.Session.CookieKey, store))

	options := sessions.Options{
		Path:     "/",
		MaxAge:   conf.Session.MaxAge,
		HttpOnly: true,
	}
	store.Options(options)

	authConf = conf

	gob.Register(conf.LoginUser)

	engine.POST(conf.Login.url, conf.Login.handle)
	engine.POST(conf.Logout.url, conf.Logout.handle)
	engine.GET(conf.LoginSuccess.url, conf.LoginSuccess.handle)
	engine.GET(conf.LoginFailed.url, conf.LoginFailed.handle)
	engine.GET(conf.UnAuthenticated.url, conf.UnAuthenticated.handle)
	engine.GET(conf.IsAuthenticated.url, conf.IsAuthenticated.handle)

}

type User interface {
	UniqueId() interface{}

	Login(c *gin.Context) (u User,isRememberMe bool, err LoginError)
}

func RequireUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !IsAuthenticated(c)  {
			c.Redirect(http.StatusMovedPermanently, authConf.UnAuthenticated.url)
			c.Abort()
			return
		}
		c.Next()
		defer context.Clear(c.Request)
	}
}

func GetUser(c *gin.Context) interface{} {
	session := sessions.Default(c)
	u := session.Get(authConf.Session.SessionKey)
	return u
}

func GetUniqueId(c *gin.Context) interface{} {
	session := sessions.Default(c)
	user := session.Get(authConf.Session.SessionKey)
	u := user.(User)
	return u.UniqueId()
}

func IsAuthenticated(c *gin.Context) bool {

	if IsLogin(c) {
		return true
	}

	err := loginByRememberMe(c)
	if err != nil {
		return false
	}

	return IsLogin(c)
}

func IsLogin(c *gin.Context) bool {
	session := sessions.Default(c)
	u := session.Get(authConf.Session.SessionKey)
	return u != nil
}

func Logout(c *gin.Context) bool {

	isRememberMeSuccess := true
	if isRememberMeEnable() {
		isRememberMeSuccess = (clearRememberMe(c) == nil)
	}
	isSuccess := (destroySession(c) == nil)

	return isSuccess && isRememberMeSuccess
}

type LoginError int

const (
	ERROR_NO_USER LoginError = (iota + 1)
	ERROR_PASSWORD_WRONG
	ERROR_NO_USER_OR_NO_PASSWORD_WRONG
	ERROR_ALREADY_LOGIN
	ERROR_ALREADY_LOGOUT
)

func destroySession(c *gin.Context) error {

	session := sessions.Default(c)
	session.Clear()
	err := session.Save()

	cookie := &http.Cookie{
		Name: authConf.Session.CookieKey,
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	http.SetCookie(c.Writer,cookie)

	return err
}

func login(c *gin.Context) {

	u,isRememberMe, err := authConf.LoginUser.Login(c)

	if err != 0 {
		c.Set("LoginError", err)
		c.Redirect(http.StatusMovedPermanently, authConf.LoginFailed.url)
		return
	}

	session := sessions.Default(c)
	session.Set(authConf.Session.SessionKey, u)
	session.Save()

	if isRememberMeEnable() && isRememberMe {
		saveRememberMe(c)
	}

	c.Redirect(http.StatusMovedPermanently, authConf.LoginSuccess.url)
}

func DefaultIsAuthenticated(c *gin.Context) {
	c.JSON(http.StatusOK, IsAuthenticated(c))
}

func DefaultUnAuthenticate(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, "user not login")
}

func DefaultLogout(c *gin.Context) {
	success := Logout(c)
	if !success {
		c.JSON(http.StatusBadRequest, "logout fail")
		return
	}
	c.JSON(http.StatusOK, "logout success")
}

func DefaultLoginSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, "login success")
}

func DefaultLoginFailed(c *gin.Context) {
	err := c.MustGet("LoginError")
	c.JSON(http.StatusOK, err)
}
