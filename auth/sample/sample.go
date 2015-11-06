package main

import (
	"github.com/gin-gonic/gin"
	"go_utils/auth"
)

type User struct {
	Id        int64
	LoginName string
	Password  string
	Salt      string
}


func (user User) UniqueId() interface{} {
	return user.Id
}

func (user User) Login() (u auth.User, err auth.LoginError) {
	return &User{
		Id:        1,
		LoginName: "小明",
	}, 0
}

func main() {
	r := gin.Default()
	auth.NewDefaultCookieSession(r, User{}, 60*30, "springcat")
	r.Use(auth.RequireUser())
	r.GET("st", st)
	r.Run(":3000")
}

func st(c *gin.Context) {

	c.JSON(200, auth.GetUser(c))
}
