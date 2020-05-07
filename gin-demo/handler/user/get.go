package user

import (
    "github.com/gin-gonic/gin"
    . "wheel/gin-demo/handler"
    "wheel/gin-demo/model"
    "wheel/gin-demo/pkg/errno"
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
