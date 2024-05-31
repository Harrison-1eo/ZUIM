package controllers

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

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

// GetCommonRooms 查找与某用户共同的房间
func GetCommonRooms(c *gin.Context) {
	// 从URL参数中解析出用户ID
	userId := c.MustGet("userID").(uint)

	friendID, err := strconv.Atoi(c.Query("friend_id"))
	if friendID == 0 || err != nil {
		respond(c, 1, "获取共同房间列表失败，请求错误", nil)
		println("// controllers/friend_controller.go 获取共同房间列表失败，请求错误 >>> err: ", friendID, err.Error())
		return
	}

	// 获取共同房间
	rooms, err := userRoomRepo.GetCommonRooms(userId, uint(friendID))
	if err != nil {
		respond(c, 1, "获取共同房间列表失败，服务器错误", nil)
		println("// controllers/friend_controller.go 获取共同房间列表失败，服务器错误 >>> err:", err.Error())
		return
	}

	type ResponseType struct {
		ID          uint   `json:"ID"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	roomsResponse := make([]ResponseType, 0)
	for _, room := range rooms {
		roomsResponse = append(roomsResponse, ResponseType{
			ID:          room.ID,
			Name:        room.Name,
			Description: room.Description,
		})
	}

	respond(c, 0, "获取共同房间列表成功", roomsResponse)
}
