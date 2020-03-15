package main

import (
	"github.com/go-xorm/xorm"
	"shortmessage/handler"
	"shortmessage/logic"
	"shortmessage/midwear"
	"shortmessage/model"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func Route(r *gin.Engine) {

	home := r.Group("/home")
	home.Use(midwear.LoginCheck())
	{
		home.POST("/add_contacts", handler.LoginHandler)
		home.POST("/contacts_list", handler.LoginHandler)
		home.POST("/del_contact", handler.LoginHandler)
		home.POST("/update_contact", handler.LoginHandler)
	}

	r.POST("/login", handler.LoginHandler)
	r.POST("/register", handler.RegisterHandler)
}

func main() {
	x, err := xorm.NewEngine("sqlite3", "data")
	if err != nil {
		panic(err)
	}

	userModel, err := model.NewUserModel(x)
	if err != nil {
		panic(err)
	}

	var (
		r = gin.New()
		l = logic.NewShortMessageLogic(userModel, "123456")
	)

	handler.RegLogic(l)
	midwear.SetSalt("123456")

	Route(r)
	_ = r.Run(":8081")
}
