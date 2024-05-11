// controllers/response_util.go

package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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

func respondWebSocket(ws *websocket.Conn, code int, msg string, data interface{}) {
	err := ws.WriteJSON(ResponseData{
		Code: code,
		Msg:  msg,
		Data: data,
	})
	if err != nil {
		println("Failed to write message")
	}
}
