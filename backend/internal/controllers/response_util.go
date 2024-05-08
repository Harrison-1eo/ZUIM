// controllers/response_util.go

package controllers

import (
	"github.com/gin-gonic/gin"
)

// ResponseData 统一响应结构
type ResponseData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 统一的响应处理函数
func respond(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(200, ResponseData{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
