package util

import (
	"github.com/gin-gonic/gin"
)

const (
	MsgSuccess = "success"
)

//ResponseBody api返回数据格式
type ResponseBody struct {
	Code    int         `json:"code,omitempty"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Msg     string      `json:"msg,omitempty"`
}

// ResponseSuccess
func ResponseSuccess(c *gin.Context, code int, data interface{}) {
	if data == nil {
		data = struct{}{}
	}
	result := ResponseBody{
		Code:    code,
		Success: true,
		Data:    data,
		Msg:     MsgSuccess,
	}
	c.SecureJSON(code, result)
}

// ResponseError
func ResponseError(c *gin.Context, code int, msg string) {
	result := ResponseBody{
		Code:    code,
		Success: false,
		Data:    struct{}{},
		Msg:     msg,
	}
	c.SecureJSON(code, result)
}
