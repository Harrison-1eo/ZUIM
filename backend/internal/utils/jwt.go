// utils/jwt.go
package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

// JWT_PREFIX JWT 前缀
const JWT_PREFIX = "Bearer"

// JWT_SECRET JWT 密钥
const JWT_SECRET = "your-secret-Key"

// CreateToken 生成JWT令牌
func CreateToken(userID uint) (string, error) {
	println("CreateToken userID: ", userID)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": float64(userID),
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // 有效期为24小时
	})
	tokenString, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return "", err
	}

	// 解析JWT令牌
	//bearerToken := strings.Split(authHeader, " ")

	return tokenString, nil
}

// ValidateToken 解析和验证JWT令牌
func ValidateToken(tokenString string) (uint, error) {
	// 检查Bearer格式
	bearerToken := strings.Split(tokenString, " ")
	if len(bearerToken) != 2 || bearerToken[0] != JWT_PREFIX {
		return 0, errors.New("Authorization header must be in 'Bearer {token}' format")
	}

	// 解析JWT令牌
	token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte("your-secret-Key"), nil
	})

	var ve *jwt.ValidationError
	if errors.As(err, &ve) {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return 0, errors.New("Malformed JWT token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return 0, errors.New("Expired or not active JWT token")
		} else {
			return 0, errors.New("Invalid JWT token")
		}
	} else if err != nil {
		return 0, errors.New("Invalid JWT token")
	} else if !token.Valid {
		return 0, errors.New("Invalid or expired JWT token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Invalid JWT token claims")
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("User ID must be a number")
	}

	return uint(userID), nil
}
