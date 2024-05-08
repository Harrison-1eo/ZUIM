// middlewares/auth_middleware.go

package middlewares

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// AuthMiddleware JWT验证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获取JWT令牌
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		// 检查Bearer格式
		println("authHeader: ", authHeader)
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header must be in 'Bearer {token}' format"})
			return
		}

		// 解析JWT令牌
		token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte("your-secret-key"), nil
		})

		var ve *jwt.ValidationError
		if errors.As(err, &ve) {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Malformed JWT token"})
				return
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Expired or not active JWT token"})
				return
			} else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid JWT token"})
				return
			}
		} else if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid JWT token"})
			return
		} else if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired JWT token"})
			return
		}

		// JWT令牌有效，提取userID并存储在请求的上下文中
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid JWT token claims"})
			return
		}

		println(" ======= claims ======== ")
		for key, value := range claims {
			fmt.Printf("%s: %v\n", key, value)
		}

		userID, ok := claims["user_id"].(float64) // JWT库默认将数字解释为float64
		println("userID: ", userID)

		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User ID must be a number"})
			return
		}

		c.Set("userID", uint(userID)) // 转换为uint类型

		//c.Set("userID", claims["user_id"])
		//println("userID: ", claims["user_id"])

		// 调用下一个中间件或处理函数
		c.Next()
	}
}
