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
		home.POST("/contacts/add", handler.AddContactsHandler)
		home.POST("/contacts/list", handler.ContactsListHandler)
		home.POST("/contact/delete", handler.DelContactHandler)
		home.POST("/contact/update", handler.UpdateContactHandler)

		home.POST("/message/send", handler.SendMessageHandler)
		home.POST("/message/view", handler.MessagePeopleViewHandler)

		home.POST("/backup/add", handler.BackupAddHandler)
		home.POST("/backup/delete", handler.BackupDeleteHandler)
		home.POST("/backup/list", handler.BackupListHandler)
		home.POST("/backup/do", handler.BackupHandler)

		home.POST("/content/add", handler.AddContentHandler)
		home.POST("/content/delete", handler.DelContentHandler)
		home.POST("/content/update", handler.UpdateContentHandler)
		home.POST("/content/list", handler.ContentsListHandler)
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

	backupDataModel, err := model.NewBackUpDataModel(x)
	if err != nil {
		panic(err)
	}

	var (
		r = gin.New()
		l = logic.NewShortMessageLogic(userModel, mailListModel, messageModel, memorandumModel, backupDataModel, "123456")
	)

	handler.RegLogic(l)
	midwear.SetSalt("123456")

	Route(r)
	_ = r.Run(":8081")
}
