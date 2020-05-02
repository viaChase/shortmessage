package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	er "shortmessage/error"
	"shortmessage/logic"
	"strconv"
)

func BackupAddHandler(context *gin.Context) {
	var req logic.BackupDataAddRequest

	if err := context.BindJSON(&req); err != nil {
		return
	}

	userId, err := strconv.ParseInt(context.Request.Header.Get("userId"), 10, 64)
	if err != nil || userId == 0 {
		context.String(http.StatusUnauthorized, er.UserNotLogin.Error())
		return
	}

	if err := shortMessageLogic.BackupAdd(&req, userId); err != nil {
		context.String(http.StatusBadGateway, err.Error())
		return
	}

	context.Status(http.StatusOK)
}

func BackupDeleteHandler(context *gin.Context) {
	var req logic.BackupDataDeleteRequest

	if err := context.BindJSON(&req); err != nil {
		return
	}

	userId, err := strconv.ParseInt(context.Request.Header.Get("userId"), 10, 64)
	if err != nil || userId == 0 {
		context.String(http.StatusUnauthorized, er.UserNotLogin.Error())
		return
	}

	if err := shortMessageLogic.BackupDelete(&req, userId); err != nil {
		context.String(http.StatusBadGateway, err.Error())
		return
	}

	context.Status(http.StatusOK)
}

func BackupListHandler(context *gin.Context) {
	var req logic.BackupDataListRequest

	if err := context.BindJSON(&req); err != nil {
		return
	}

	userId, err := strconv.ParseInt(context.Request.Header.Get("userId"), 10, 64)
	if err != nil || userId == 0 {
		context.String(http.StatusUnauthorized, er.UserNotLogin.Error())
		return
	}

	if resp, err := shortMessageLogic.BackupList(&req, userId); err != nil {
		context.String(http.StatusBadGateway, err.Error())
		return
	} else {
		context.JSON(http.StatusOK, resp)
	}

}

func BackupHandler(context *gin.Context) {
	var req logic.BackupRequest

	if err := context.BindJSON(&req); err != nil {
		return
	}

	userId, err := strconv.ParseInt(context.Request.Header.Get("userId"), 10, 64)
	if err != nil || userId == 0 {
		context.String(http.StatusUnauthorized, er.UserNotLogin.Error())
		return
	}

	if err := shortMessageLogic.BackUp(&req, userId); err != nil {
		context.String(http.StatusBadGateway, err.Error())
		return
	}

	context.Status(http.StatusOK)
}
