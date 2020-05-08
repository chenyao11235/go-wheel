package user

import (
	"github.com/gin-gonic/gin"
	"strconv"
	. "wheel/gin-demo/handler"
	"wheel/gin-demo/logger"
	"wheel/gin-demo/model"
	"wheel/gin-demo/pkg/errno"
)

func Update(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))

	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u.Id = uint64((userId))
	if err := u.Validation(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	if err := u.Update(); err != nil {
		logger.Log.Error(err)
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
