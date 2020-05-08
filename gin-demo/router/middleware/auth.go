package middleware

import (
	"github.com/gin-gonic/gin"
	"wheel/gin-demo/handler"
	"wheel/gin-demo/pkg/errno"
	"wheel/gin-demo/pkg/token"
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
