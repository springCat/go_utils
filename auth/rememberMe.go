package auth

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	gorillaSession "github.com/gorilla/sessions"
	"errors"
	"log"
)

const (
	RMDefaultKey = "springCatRememberMe"
)

type RememberMeConf struct {
	CookieKey string
	MaxAge    int
	Secret    string
}

var rememberMeConf RememberMeConf

var cookieStore sessions.CookieStore

func NewRememberMe(engine *gin.Engine, cookieKey string, maxAge int, secret string) {

	rememberMeConf = RememberMeConf{
		CookieKey: cookieKey,
		MaxAge:    maxAge,
		Secret:    secret,
	}

	options := sessions.Options{
		Path:     "/",
		MaxAge:   rememberMeConf.MaxAge,
		HttpOnly: true,
	}
	cookieStore = sessions.NewCookieStore([]byte(rememberMeConf.Secret))
	cookieStore.Options(options)
}

func isRememberMeEnable() bool {
	return cookieStore != nil
}

func getRMSession(c *gin.Context) *gorillaSession.Session {
	RMSession, err := cookieStore.Get(c.Request, rememberMeConf.CookieKey)
	if err != nil {
		log.Println(err)
	}
	return RMSession
}

func getRMSessionValue(c *gin.Context, key string) interface{} {
	RMSession := getRMSession(c)
	return RMSession.Values[key]
}

func setRMSessionValue(c *gin.Context, key string, value interface{}) *gorillaSession.Session {
	RMSession := getRMSession(c)
	RMSession.Values[key] = value
	return RMSession
}

func saveRMSession(c *gin.Context, session *gorillaSession.Session) error {
	return cookieStore.Save(c.Request, c.Writer, session)
}

func destroyRMSession(c *gin.Context) error {
	RMSession := getRMSession(c)
	RMSession.Options = &gorillaSession.Options{
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	return saveRMSession(c, RMSession)
}

func saveRememberMe(c *gin.Context) error {
	session := sessions.Default(c)
	loginUser := session.Get(authConf.Session.SessionKey)

	RMSession := setRMSessionValue(c, authConf.Session.SessionKey, loginUser)
	return saveRMSession(c, RMSession)
}

func loginByRememberMe(c *gin.Context) error {

	loginUser := getRMSessionValue(c, authConf.Session.SessionKey)

	if loginUser == nil {
		return errors.New("no rememberMe")
	}

	session := sessions.Default(c)
	session.Set(authConf.Session.SessionKey, loginUser)
	return session.Save()
}

func clearRememberMe(c *gin.Context) error {
	return destroyRMSession(c)
}
