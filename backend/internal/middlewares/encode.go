package middlewares

import (
	"backend/internal/controllers"
	"bytes"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type decryptBody struct {
	Length   int    `json:"length"`
	Position int    `json:"position"`
	Mac      string `json:"mac"`
	Data     string `json:"data"`
}

func CipherDecryptMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Body == nil || c.Request.ContentLength == 0 || c.Request.Method == "GET" {
			// 如果Body为空，直接跳过
		} else {
			// 如果访问路径中包含upload，则不解码
			if strings.Contains(c.Request.URL.Path, "upload") {
				// 不做处理
				//println("No need to decode: ", c.Request.URL.Path)
			} else {
				userID := c.MustGet("userID").(uint)

				// 根据decryptBody结构体解析数据
				var db decryptBody
				err := c.ShouldBindJSON(&db)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to decode request body: " + err.Error()})
					return
				}

				//println("Before Decoded data: ", db.Data)

				// 对decodedMsg进行解密
				uf := controllers.UserCipherFrontends[userID]
				if uf == nil {
					println("Failed to get user cipher frontend: ", userID)
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to get user cipher frontend"})
					return
				}
				decodedBody, err := uf.Decrypt(db.Data, db.Position) // 再进行解密
				if err != nil {
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to decode request body: " + err.Error()})
					return
				}

				//println("Decoded body: ", decodedBody)

				// 重置Body
				c.Request.Body = io.NopCloser(bytes.NewBuffer([]byte(decodedBody)))

			}
		}

		c.Next()
	}
}
