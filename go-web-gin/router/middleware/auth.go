package middleware

import (
	"github.com/gin-gonic/gin"
    "wheel/go-web-gin/handler"
    "wheel/go-web-gin/pkg/errno"
    "wheel/go-web-gin/pkg/token"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
