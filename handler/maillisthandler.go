package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	er "shortmessage/error"
	"shortmessage/logic"
	"strconv"
)

func AddContactsHandler(context *gin.Context) {
	var req logic.AddContactsRequest

	if err := context.ShouldBindJSON(&req); err != nil {
		return
	}

	userId, err := strconv.ParseInt(context.Request.Header.Get("userId"), 10, 64)
	if err != nil || userId == 0 {
		context.String(http.StatusUnauthorized, er.UserNotLogin.Error())
		return
	}

	if err := shortMessageLogic.AddContacts(&req, userId); err != nil {
		context.String(http.StatusBadGateway, err.Error())
		return
	}

	context.Status(http.StatusOK)
}

func ContactsListHandler(context *gin.Context) {
	var req logic.ContactsListRequest

	if err := context.ShouldBindJSON(&req); err != nil {
		return
	}

	userId, err := strconv.ParseInt(context.Request.Header.Get("userId"), 10, 64)
	if err != nil || userId == 0 {
		context.String(http.StatusUnauthorized, er.UserNotLogin.Error())
		return
	}

	resp, err := shortMessageLogic.ContactsList(&req, userId)
	if err != nil {
		context.String(http.StatusBadGateway, err.Error())
		return
	}

	context.JSON(http.StatusOK, resp)
}

func DelContactHandler(context *gin.Context) {
	var req logic.DeleteContactRequest

	if err := context.ShouldBindJSON(&req); err != nil {
		return
	}

	userId, err := strconv.ParseInt(context.Request.Header.Get("userId"), 10, 64)
	if err != nil || userId == 0 {
		context.String(http.StatusUnauthorized, er.UserNotLogin.Error())
		return
	}

	if err := shortMessageLogic.DelContact(&req, userId); err != nil {
		context.String(http.StatusBadGateway, err.Error())
		return
	}

	context.Status(http.StatusOK)
}

func UpdateContactHandler(context *gin.Context) {
	var req logic.UpdateContactsRequest

	if err := context.ShouldBindJSON(&req); err != nil {
		return
	}

	userId, err := strconv.ParseInt(context.Request.Header.Get("userId"), 10, 64)
	if err != nil || userId == 0 {
		context.String(http.StatusUnauthorized, er.UserNotLogin.Error())
		return
	}

	if err := shortMessageLogic.UpdateContact(&req, userId); err != nil {
		context.String(http.StatusBadGateway, err.Error())
		return
	}

	context.Status(http.StatusOK)
}
