package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shortmessage/logic"
	"strconv"
)

func AddContactsHandler(context *gin.Context) {
	var req logic.AddContactsRequest

	if err := context.BindQuery(&req); err != nil {
		return
	}

	userId, err := strconv.ParseInt(context.Request.Header.Get("userId"), 10, 64)
	if err != nil || userId == 0 {
		_ = context.AbortWithError(http.StatusForbidden, err)
		return
	}

	if err := shortMessageLogic.AddContacts(&req, userId); err != nil {
		_ = context.AbortWithError(http.StatusForbidden, err)
		return
	}

	context.Status(http.StatusOK)
}

func ContactsListHanlder(context *gin.Context) {
	var req logic.ContactsListRequest

	if err := context.BindQuery(&req); err != nil {
		return
	}

	userId, err := strconv.ParseInt(context.Request.Header.Get("userId"), 10, 64)
	if err != nil || userId == 0 {
		_ = context.AbortWithError(http.StatusForbidden, err)
		return
	}

	resp, err := shortMessageLogic.ContactsList(&req, userId)
	if err != nil {
		_ = context.AbortWithError(http.StatusForbidden, err)
		return
	}

	context.JSON(http.StatusOK, resp)
}

func DelContactHandler(context *gin.Context) {
	var req logic.DeleteContactRequest

	if err := context.BindQuery(&req); err != nil {
		return
	}

	userId, err := strconv.ParseInt(context.Request.Header.Get("userId"), 10, 64)
	if err != nil || userId == 0 {
		_ = context.AbortWithError(http.StatusForbidden, err)
		return
	}

	if err := shortMessageLogic.DelContact(&req, userId); err != nil {
		_ = context.AbortWithError(http.StatusForbidden, err)
		return
	}

	context.Status(http.StatusOK)
}

func UpdateContactHandler(context *gin.Context) {
	var req logic.UpdateContactsRequest

	if err := context.BindQuery(&req); err != nil {
		return
	}

	userId, err := strconv.ParseInt(context.Request.Header.Get("userId"), 10, 64)
	if err != nil || userId == 0 {
		_ = context.AbortWithError(http.StatusForbidden, err)
		return
	}

	if err := shortMessageLogic.UpdateContact(&req, userId); err != nil {
		_ = context.AbortWithError(http.StatusForbidden, err)
		return
	}

	context.Status(http.StatusOK)
}
