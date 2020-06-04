package user

import (
	"github.com/gin-gonic/gin"
    . "wheel/go-web-gin/handler"
    "wheel/go-web-gin/model"
    "wheel/go-web-gin/pkg/errno"
)

func Get(c *gin.Context) {
	username := c.Param("username")
	user, err := model.GetUser(username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	SendResponse(c, nil, user)
}
