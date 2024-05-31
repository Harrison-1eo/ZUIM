// controllers/response_util.go

package controllers

import (
	"backend/internal/models"
	"backend/internal/utils"
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
		Mac:      utils.ZUCMac32(UserEncryptKey[userID][0], string(jsonStr)),
	}

	c.JSON(200, ResponseData{
		Code:   code,
		Msg:    msg,
		Data:   empty{},
		EnData: db,
	})
}

func respondWebSocket(userID uint, ws *websocket.Conn, code int, msg string, data interface{}) {
	if code != 0 {
		err := ws.WriteJSON(ResponseData{
			Code:   code,
			Msg:    msg,
			Data:   data,
			EnData: empty{},
		})
		if err != nil {
			println("Failed to write message")
		}
	} else {
		// 首先将data转换为models.MessageResponseBody类型
		msgData, ok := data.(models.MessageResponseBody)
		if !ok {
			println("Failed to convert data to models.MessageResponseBody")
			return
		}
		decryptData, pos, err := UserCipherWebsocketsBackends[userID].Encrypt([]byte(msgData.Content))
		if err != nil {
			println("Failed to encrypt message data: ", err.Error())
			return
		}

		msgData.Content = ""

		db := decryptBody{
			Length:   len([]byte(msgData.Content)),
			Position: pos,
			Mac:      utils.ZUCMac32(UserEncryptKey[userID][2], msgData.Content),
			Data:     decryptData,
		}

		err = ws.WriteJSON(ResponseData{
			Code:   code,
			Msg:    msg,
			Data:   msgData,
			EnData: db,
		})

		if err != nil {
			println("Failed to write message")
		}
	}

}
