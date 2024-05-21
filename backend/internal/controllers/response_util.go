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
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
	EnData interface{} `json:"en_data"`
}

type decryptBody struct {
	Length   int    `json:"length"`
	Position int    `json:"position"`
	Mac      string `json:"mac"`
	Data     string `json:"data"`
}

type empty struct{}

func respondWithoutEncrypt(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(200, ResponseData{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

// 统一的响应处理函数
func respond(c *gin.Context, code int, msg string, data interface{}) {
	userID := c.MustGet("userID").(uint)
	// 首先将数据转为JSON字符串
	jsonStr, _ := json.Marshal(data)

	// 首先进行加密
	dataByteLength := len(jsonStr)
	decryptedData, pos, err := UserCipherBackends[userID].Encrypt(jsonStr)
	if err != nil {
		respond(c, 500, "Failed to encrypt data: "+err.Error(), nil)
		return
	}

	db := decryptBody{
		Position: pos,
		Length:   dataByteLength,
		Data:     decryptedData,
		Mac:      "",
	}

	c.JSON(200, ResponseData{
		Code:   code,
		Msg:    msg,
		Data:   empty{},
		EnData: db,
	})
}

func respondWebSocket(ws *websocket.Conn, code int, msg string, data interface{}) {
	dataBase64 := base64Encode(data)
	err := ws.WriteJSON(ResponseData{
		Code: code,
		Msg:  msg,
		Data: dataBase64,
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
