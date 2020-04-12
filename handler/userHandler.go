package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shortmessage/logic"
)

func LoginHandler(context *gin.Context) {
	var req logic.LoginRequest

	if err := context.BindJSON(&req); err != nil {
		return
	}

	//业务逻辑
	userId, jwt, resp, err := shortMessageLogic.Login(&req)
	if err != nil {
		context.String(http.StatusUnauthorized, err.Error())
		return
	}

	context.Header("userId", fmt.Sprintf("%v", userId))
	context.Header("jwt", jwt)
	context.JSON(http.StatusOK, resp)
}

func RegisterHandler(context *gin.Context) {
	var req logic.RegisterRequest

	if err := context.BindJSON(&req); err != nil {
		return
	}

	resp, err := shortMessageLogic.Register(&req)
	if err != nil {
		context.String(http.StatusUnauthorized, err.Error())
		return
	}

	context.JSON(http.StatusOK, resp)
}
