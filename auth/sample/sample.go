package main

import (
	"github.com/gin-gonic/gin"
	"github.com/springCat/go_utils/auth"
)

type User struct {
	Id        int64
	LoginName string
}


func (user User) UniqueId() interface{} {
	return user.Id
}

func (user User) Login(c *gin.Context) (u auth.User,isRememberMe bool, err auth.LoginError) {
	return &User{
		Id:        1,
		LoginName: "小明",
	},false, 0
}

func main() {
	r := gin.Default()
	auth.NewDefaultCookieSession(r, User{}, 60*30, "springcat")
	auth.NewRememberMe(r,"RememberMe",30*24*60*60,"springcat")
	r.Use(auth.RequireUser())
	r.GET("st", st)
	r.Run(":3000")
}

func st(c *gin.Context) {

	c.JSON(200, auth.GetUser(c))
}
