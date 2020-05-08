package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/willf/pad"
	"io/ioutil"
	"regexp"
	"time"
	"wheel/gin-demo/handler"
	"wheel/gin-demo/logger"
	"wheel/gin-demo/pkg/errno"
)

type bodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)

}

func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now().UTC()
		path := c.Request.URL.Path

		// 只对登陆和用户操作的api记录日志
		reg := regexp.MustCompile("/v1/user|/login")
		if !reg.MatchString(path) {
			return
		}

		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}

		// 因为request body 在读取之后会被置空，所以需要重新存储
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		method := c.Request.Method
		ip := c.ClientIP()

		blw := &bodyWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}

		c.Writer = blw
		c.Next()

		end := time.Now().UTC()
		latency := end.Sub(start)

		code, message := -1, ""
		var response handler.Response
		if err := json.Unmarshal(blw.body.Bytes(), &response); err != nil {
			logger.Log.Errorf("response body can not unmarshal to model.Response struct, body: `%s` %s", blw.body.Bytes(), err)
			code = errno.InternalServerError.Code
			message = err.Error()
		} else {
			code = response.Code
			message = response.Message
		}
		logger.Log.Infof("%-13s | %-12s | %s %s | {code: %d, message: %s}", latency, ip, pad.Right(method, 5, ""), path, code, message)
	}
}
