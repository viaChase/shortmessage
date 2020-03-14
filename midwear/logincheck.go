package midwear

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shortmessage/common"
)

var (
	salt = ""
)

func SetSalt(salts string) {
	salt = salts
}

//这个拿来验证是否登入
func LoginCheck() gin.HandlerFunc {

	return func(c *gin.Context) {
		userId := c.Request.Header.Get("userId")
		jwt := c.Request.Header.Get("jwt")

		if userId == "" || jwt == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if jwt != common.Md5(fmt.Sprintf("%v-%v", userId, salt)) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
