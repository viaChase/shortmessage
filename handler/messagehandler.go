package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
		_ = context.AbortWithError(http.StatusForbidden, err)
		return
	}

	err = shortMessageLogic.SendMessage(&req, userId)
	if err != nil {
		_ = context.AbortWithError(http.StatusForbidden, err)
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
		_ = context.AbortWithError(http.StatusForbidden, err)
		return
	}

	resp, err := shortMessageLogic.MessagePeopleView(&req, userId)
	if err != nil {
		_ = context.AbortWithError(http.StatusForbidden, err)
		return
	}

	context.JSON(http.StatusOK, resp)
}
