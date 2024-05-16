package middlewares

import (
	"bytes"
	"encoding/base64"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Base64DecodeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Body == nil {
			// 如果Body为空，直接跳过
		} else {
			// 读取Body
			buf, err := io.ReadAll(c.Request.Body)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
				return
			}

			// 解码Body
			println("Send Before decode: ", string(buf))
			println(string(buf) == "eyJ1c2VybmFtZSI6IuaIkeeahCIsInBhc3N3b3JkIjoiYWRtaW4ifQ==")
			decodedBody, err := base64.StdEncoding.DecodeString(string(buf)) // 改用StdEncoding
			println("Send After decode: ", string(decodedBody))
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to decode request body: " + err.Error()})
				return
			}

			// 重置Body
			c.Request.Body = io.NopCloser(bytes.NewBuffer(decodedBody))
		}

		c.Next()
	}
}
