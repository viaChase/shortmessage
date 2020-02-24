package main

import (
	"github.com/go-xorm/xorm"
	"shortmessage/handler"
	"shortmessage/logic"
	"shortmessage/model"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func Route(r *gin.Engine) {
	r.GET("/ping", handler.TestHandler)
	r.POST("/login", handler.LoginHandler)
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
		l = logic.NewShortMessageLogic(userModel)
	)

	handler.RegLogic(l)
	Route(r)
	_ = r.Run(":8081")
}
