// controllers/user_controller.go

package controllers

import (
	"IMTest/internal/models"
	"IMTest/internal/repositories"
	"IMTest/internal/utils"
	"github.com/gin-gonic/gin"
)

var userRepository = repositories.NewUserRepository()

// Register 注册用户
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		respond(c, 1, "注册失败，请求错误", nil)
		println("// controllers/user_controller.go 注册失败，请求错误 >>> err:", err.Error())
		return
	}

	// 检查用户名是否已存在
	if userRepository.UserExists(user.Username) {
		respond(c, 1, "注册失败，用户名已存在", nil)
		println("// controllers/user_controller.go 注册失败，用户名已存在 >>> err:")
		return
	}

	// 创建用户
	newUser, err := userRepository.CreateUser(user)
	if err != nil {
		respond(c, 1, "注册失败，服务器错误", nil)
		println("// controllers/user_controller.go 注册失败，服务器错误 >>> err:", err.Error())
		return
	}

	respond(c, 0, "注册成功", newUser)
}

// 用户登录
func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		respond(c, 1, "登录失败，请求错误", nil)
		println("// controllers/user_controller.go 登录失败，请求错误 >>> err:", err.Error())
		return
	}

	// 验证用户登录信息
	authenticatedUser, err := userRepository.AuthenticateUser(user.Username, user.Password)
	if err != nil {
		respond(c, 1, "登录失败，用户名或密码错误", nil)
		println("// controllers/user_controller.go 登录失败，用户名或密码错误 >>> err:", err.Error())
		return
	}

	// 生成JWT令牌
	token, err := utils.CreateToken(authenticatedUser.ID)
	if err != nil {
		respond(c, 1, "登录失败，服务器错误", nil)
		println("// controllers/user_controller.go 登录失败，服务器错误 >>> err:", err.Error())
		return
	}

	respond(c, 0, "登录成功", gin.H{"user": authenticatedUser, "token": token})
}
