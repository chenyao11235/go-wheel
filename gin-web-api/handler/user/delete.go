package user

import (
	"github.com/gin-gonic/gin"
	"strconv"
    . "wheel/go-web-gin/handler"
    "wheel/go-web-gin/model"
    "wheel/go-web-gin/pkg/errno"
)

func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.Delete(uint64(userId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	SendResponse(c, nil, nil)
}
