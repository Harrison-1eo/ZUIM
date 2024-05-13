package controllers

import "github.com/gin-gonic/gin"

func GetFriends(c *gin.Context) {
	// 测试用途，返回所有用户列表
	userInfos, err := userRepository.GetAllUsers()
	if err != nil {
		respond(c, 1, "获取用户列表失败", nil)
	}

	type ResponseType struct {
		ID       uint   `json:"ID"`
		Username string `json:"username"`
		Avatar   string `json:"avatar"`
		NikeName string `json:"nike_name"`
	}

	users := make([]ResponseType, 0)
	for _, userInfo := range userInfos {
		users = append(users, ResponseType{
			ID:       userInfo.UserID,
			Username: userInfo.Username,
			Avatar:   userInfo.Avatar,
			NikeName: userInfo.NikeName,
		})
	}

	respond(c, 0, "获取用户列表成功", users)

}
