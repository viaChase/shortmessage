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
		home.POST("/add_contacts", handler.AddContactsHandler)
		home.POST("/contacts_list", handler.ContactsListHandler)
		home.POST("/del_contact", handler.DelContactHandler)
		home.POST("/update_contact", handler.UpdateContactHandler)
		home.POST("/send_message", handler.SendMessageHandler)
		home.POST("/message_view", handler.MessagePeopleViewHandler)
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

	mailListModel, err := model.NewMailListModel(x)
	if err != nil {
		panic(err)
	}

	messageModel, err := model.NewMessageModel(x)
	if err != nil {
		panic(err)
	}

	memorandumModel, err := model.NewMemorandumModel(x)
	if err != nil {
		panic(err)
	}

	var (
		r = gin.New()
		l = logic.NewShortMessageLogic(userModel, mailListModel, messageModel, memorandumModel, "123456")
	)

	handler.RegLogic(l)
	midwear.SetSalt("123456")

	Route(r)
	_ = r.Run(":8081")
}
