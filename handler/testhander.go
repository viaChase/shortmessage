package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shortmessage/logic"
)

func TestHandler(context *gin.Context) {
	var req logic.TestRequest

	if err := context.BindQuery(&req); err != nil {
		return
	}

	//业务逻辑
	resp, err := shortMessageLogic.Test(&req)
	if err != nil {
		return
	}

	context.JSON(http.StatusOK, resp)
}

func LoginHandler(context *gin.Context) {
	var req logic.LoginRequest

	if err := context.BindQuery(&req); err != nil {
		return
	}

	//业务逻辑
	resp, err := shortMessageLogic.Login(&req)
	if err != nil {
		return
	}

	context.JSON(http.StatusOK, resp)
}
