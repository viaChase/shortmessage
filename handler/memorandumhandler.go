package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shortmessage/logic"
	"strconv"
)

func AddContentHandler(context *gin.Context) {
	var req logic.AddContentRequest

	if err := context.BindQuery(&req); err != nil {
		return
	}

	userId, err := strconv.ParseInt(context.Request.Header.Get("userId"), 10, 64)
	if err != nil || userId == 0 {
		_ = context.AbortWithError(http.StatusForbidden, err)
		return
	}

	if err := shortMessageLogic.AddContent(&req, userId); err != nil {
		_ = context.AbortWithError(http.StatusForbidden, err)
		return
	}

	context.Status(http.StatusOK)
}

func DelContentHandler(context *gin.Context) {
	var req logic.DeleteContentRequest

	if err := context.BindQuery(&req); err != nil {
		return
	}

	userId, err := strconv.ParseInt(context.Request.Header.Get("userId"), 10, 64)
	if err != nil || userId == 0 {
		_ = context.AbortWithError(http.StatusForbidden, err)
		return
	}

	if err := shortMessageLogic.DelContent(&req, userId); err != nil {
		_ = context.AbortWithError(http.StatusForbidden, err)
		return
	}

	context.Status(http.StatusOK)
}

func UpdateContentHandler(context *gin.Context) {
	var req logic.UpdateContentRequest

	if err := context.BindQuery(&req); err != nil {
		return
	}

	userId, err := strconv.ParseInt(context.Request.Header.Get("userId"), 10, 64)
	if err != nil || userId == 0 {
		_ = context.AbortWithError(http.StatusForbidden, err)
		return
	}

	if err := shortMessageLogic.Update(&req, userId); err != nil {
		_ = context.AbortWithError(http.StatusForbidden, err)
		return
	}

	context.Status(http.StatusOK)
}

func ContentsListHanlder(context *gin.Context) {
	var req logic.ContentListRequest

	if err := context.BindQuery(&req); err != nil {
		return
	}

	userId, err := strconv.ParseInt(context.Request.Header.Get("userId"), 10, 64)
	if err != nil || userId == 0 {
		_ = context.AbortWithError(http.StatusForbidden, err)
		return
	}

	resp, err := shortMessageLogic.ContentList(&req, userId)
	if err != nil {
		_ = context.AbortWithError(http.StatusForbidden, err)
		return
	}

	context.JSON(http.StatusOK, resp)
}
