package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shortmessage/logic"
)

func LoginHandler(context *gin.Context) {
	var req logic.LoginRequest

	if err := context.BindQuery(&req); err != nil {
		return
	}

	//业务逻辑
	userId, jwt, err := shortMessageLogic.Login(&req)
	if err != nil {
		return
	}

	context.Header("userId", fmt.Sprintf("%v", userId))
	context.Header("jwt", jwt)
}

func RegisterHandler(context *gin.Context) {
	var req logic.RegisterRequest

	if err := context.BindQuery(&req); err != nil {
		return
	}

	resp, err := shortMessageLogic.Register(&req)
	if err != nil {
		_ = context.AbortWithError(http.StatusForbidden, err)
		return
	}

	context.JSON(http.StatusOK, resp)
}
