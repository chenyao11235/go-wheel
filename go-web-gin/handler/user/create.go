package user

import (
	"github.com/gin-gonic/gin"
    . "wheel/go-web-gin/handler"
    "wheel/go-web-gin/logger"
    "wheel/go-web-gin/model"
    "wheel/go-web-gin/pkg/errno"
    "wheel/go-web-gin/util"
)

func Create(c *gin.Context) {
	logger.Log.Infof("User create function called, X-Request-Id is %s", util.GetReqID(c))

	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	if err := u.Validation(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	if err := u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	resp := CreateResponse{
		Username: r.Username,
	}

	SendResponse(c, nil, resp)
}
