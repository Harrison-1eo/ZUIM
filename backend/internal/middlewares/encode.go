package middlewares

import (
	"bytes"
	"encoding/base64"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Base64DecodeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Body == nil {
			// 如果Body为空，直接跳过
		} else {
			// 如果访问路径中包含upload，则不解码
			if strings.Contains(c.Request.URL.Path, "upload") {
				// 不做处理
				println("No need to decode: ", c.Request.URL.Path)
			} else {
				// 读取Body，解码Body
				buf, err := io.ReadAll(c.Request.Body)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
					return
				}

				println("Send Before decode: ", string(buf))
				decodedBody, err := base64.StdEncoding.DecodeString(string(buf)) // 改用StdEncoding
				println("Send After decode: ", string(decodedBody))
				if err != nil {
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to decode request body: " + err.Error()})
					return
				}

				// 重置Body
				c.Request.Body = io.NopCloser(bytes.NewBuffer(decodedBody))
			}
		}

		c.Next()
	}
}
