package user

import (
	"github.com/gin-gonic/gin"
	"strconv"
	. "wheel/gin-demo/handler"
	"wheel/gin-demo/model"
	"wheel/gin-demo/pkg/errno"
)

func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.Delete(uint64(userId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	SendResponse(c, nil, nil)
}
