// controllers/response_util.go

package controllers

import (
	"encoding/base64"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// ResponseData 统一响应结构
type ResponseData struct {
	Code     int         `json:"code"`
	Msg      string      `json:"msg"`
	Length   int         `json:"length"`
	Position int         `json:"position"`
	Data     interface{} `json:"data"`
}

// 统一的响应处理函数
func respond(c *gin.Context, code int, msg string, data interface{}) {
	dataBase64 := base64Encode(data)
	c.JSON(200, ResponseData{
		Code:     code,
		Msg:      msg,
		Length:   base64Length(dataBase64),
		Position: 0,
		Data:     dataBase64,
	})
}

func respondWebSocket(ws *websocket.Conn, code int, msg string, data interface{}) {
	dataBase64 := base64Encode(data)
	err := ws.WriteJSON(ResponseData{
		Code:     code,
		Msg:      msg,
		Length:   base64Length(dataBase64),
		Position: 0,
		Data:     dataBase64,
	})
	if err != nil {
		println("Failed to write message")
	}
}

func base64Encode(data interface{}) string {
	// 先将数据转为json字符串
	jsonStr, _ := json.Marshal(data)
	// 再进行base64编码
	return base64.StdEncoding.EncodeToString(jsonStr)
}

func base64Length(data string) int {
	return len(data)
}
