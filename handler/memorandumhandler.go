package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	er "shortmessage/error"
	"shortmessage/logic"
	"strconv"
)

func AddContentHandler(context *gin.Context) {
	var req logic.AddContentRequest

	if err := context.BindJSON(&req); err != nil {
		return
	}

	userId, err := strconv.ParseInt(context.Request.Header.Get("userId"), 10, 64)
	if err != nil || userId == 0 {
		context.String(http.StatusUnauthorized, er.UserNotLogin.Error())
		return
	}

	if err := shortMessageLogic.AddContent(&req, userId); err != nil {
		context.String(http.StatusBadGateway, err.Error())
		return
	}

	context.Status(http.StatusOK)
}

func DelContentHandler(context *gin.Context) {
	var req logic.DeleteContentRequest

	if err := context.BindJSON(&req); err != nil {
		return
	}

	userId, err := strconv.ParseInt(context.Request.Header.Get("userId"), 10, 64)
	if err != nil || userId == 0 {
		context.String(http.StatusUnauthorized, er.UserNotLogin.Error())
		return
	}

	if err := shortMessageLogic.DelContent(&req, userId); err != nil {
		context.String(http.StatusBadGateway, err.Error())
		return
	}

	context.Status(http.StatusOK)
}

func UpdateContentHandler(context *gin.Context) {
	var req logic.UpdateContentRequest

	if err := context.BindJSON(&req); err != nil {
		return
	}

	userId, err := strconv.ParseInt(context.Request.Header.Get("userId"), 10, 64)
	if err != nil || userId == 0 {
		context.String(http.StatusUnauthorized, er.UserNotLogin.Error())
		return
	}

	if err := shortMessageLogic.Update(&req, userId); err != nil {
		context.String(http.StatusBadGateway, err.Error())
		return
	}

	context.Status(http.StatusOK)
}

func ContentsListHandler(context *gin.Context) {
	var req logic.ContentListRequest

	if err := context.BindJSON(&req); err != nil {
		return
	}

	userId, err := strconv.ParseInt(context.Request.Header.Get("userId"), 10, 64)
	if err != nil || userId == 0 {
		context.String(http.StatusUnauthorized, er.UserNotLogin.Error())
		return
	}

	resp, err := shortMessageLogic.ContentList(&req, userId)
	if err != nil {
		context.String(http.StatusBadGateway, err.Error())
		return
	}

	context.JSON(http.StatusOK, resp)
}
