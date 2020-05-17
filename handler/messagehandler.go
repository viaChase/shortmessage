package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	er "shortmessage/error"
	"shortmessage/logic"
	"strconv"
)

func SendMessageHandler(context *gin.Context) {
	var req logic.SendMessageRequest

	if err := context.BindJSON(&req); err != nil {
		return
	}

	userId, err := strconv.ParseInt(context.Request.Header.Get("userId"), 10, 64)
	if err != nil || userId == 0 {
		context.String(http.StatusUnauthorized, er.UserNotLogin.Error())
		return
	}

	err = shortMessageLogic.SendMessage(&req, userId)
	if err != nil {
		context.String(http.StatusBadGateway, err.Error())
		return
	}

	context.Status(http.StatusOK)
}

func MessagePeopleViewHandler(context *gin.Context) {
	var req logic.MessagePeopleViewRequest

	if err := context.BindJSON(&req); err != nil {
		return
	}

	userId, err := strconv.ParseInt(context.Request.Header.Get("userId"), 10, 64)
	if err != nil || userId == 0 {
		context.String(http.StatusUnauthorized, er.UserNotLogin.Error())
		return
	}

	resp, err := shortMessageLogic.MessagePeopleView(&req, userId)
	if err != nil {
		context.String(http.StatusBadGateway, err.Error())
		return
	}

	context.JSON(http.StatusOK, resp)
}

func MessageReadHandler(context *gin.Context) {
	var req logic.MessageReadRequest

	if err := context.BindJSON(&req); err != nil {
		return
	}

	userId, err := strconv.ParseInt(context.Request.Header.Get("userId"), 10, 64)
	if err != nil || userId == 0 {
		fmt.Println("aa")
		context.String(http.StatusUnauthorized, er.UserNotLogin.Error())
		return
	}

	shortMessageLogic.MessageRead(&req, userId)
	context.Status(http.StatusOK)
}

func MessageDeleteHandler(context *gin.Context) {
	var req logic.MessageDeleteRequest

	if err := context.BindJSON(&req); err != nil {
		return
	}

	userId, err := strconv.ParseInt(context.Request.Header.Get("userId"), 10, 64)
	if err != nil || userId == 0 {
		context.String(http.StatusUnauthorized, er.UserNotLogin.Error())
		return
	}

	shortMessageLogic.MessageDelete(&req, userId)
	context.Status(http.StatusOK)
}
