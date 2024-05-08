// utils/jwt.go
package utils

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// CreateToken 生成JWT令牌
func CreateToken(userID uint) (string, error) {
	println("CreateToken userID: ", userID)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": float64(userID),
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // 有效期为24小时
	})
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}

	// 解析JWT令牌
	//bearerToken := strings.Split(authHeader, " ")

	tokenCheck, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte("your-secret-key"), nil
	})

	claims, _ := tokenCheck.Claims.(jwt.MapClaims)

	println(" ======= jwt claims ======== ")
	for key, value := range claims {
		fmt.Printf("%s: %v\n", key, value)
	}

	return tokenString, nil
}

// ValidateToken 解析和验证JWT令牌
func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method in token")
		}
		return []byte("your-secret-key"), nil
	})
}
