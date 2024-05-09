// controllers/room_controller.go

package controllers

import (
	"backend/internal/models"
	"backend/internal/repositories"
	"github.com/gin-gonic/gin"
)

var roomRepo *repositories.RoomRepository = repositories.NewRoomRepository()
var userRoomRepo *repositories.UserRoomRepository = repositories.NewUserRoomRepository()

// CreateRoom 创建聊天室
func CreateRoom(c *gin.Context) {
	type RequestBody struct {
		RoomName string `json:"name"`
		RoomInfo string `json:"description"`
	}

	// 从请求体中解析出聊天室信息
	var body RequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		respond(c, 1, "创建聊天室失败，请求错误", nil)
		println("// controllers/room_controller.go 创建聊天室失败，请求错误 >>> err:", err.Error())
		return
	}

	if body.RoomName == "" {
		respond(c, 1, "创建聊天室失败，聊天室名称不能为空", nil)
		println("// controllers/room_controller.go 创建聊天室失败，聊天室名称不能为空 >>> err:")
		return
	}

	// 创建Room对象
	room := models.Room{
		Name:        body.RoomName,
		Description: body.RoomInfo,
	}

	newRoom, err := roomRepo.CreateRoom(room)
	if err != nil {
		respond(c, 1, "创建聊天室失败，服务器错误1", nil)
		println("// controllers/room_controller.go 创建聊天室失败，服务器错误1 >>> err:", err.Error())
		return
	}

	// 将当前用户添加到聊天室
	userId := c.MustGet("userID").(uint)
	if err := userRoomRepo.AddUserToRoom(userId, newRoom.ID, "owner"); err != nil {
		respond(c, 1, "创建聊天室失败，服务器错误2", nil)
		println("// controllers/room_controller.go 创建聊天室失败，服务器错误2 >>> err:", err.Error())
		return
	}

	respond(c, 0, "创建聊天室成功", newRoom)
}

// GetRoomList 获取我的聊天室列表
func GetRoomList(c *gin.Context) {
	userId := c.MustGet("userID").(uint)
	rooms, err := userRoomRepo.GetRoomsByUserID(userId)
	if err != nil {
		respond(c, 1, "获取聊天室列表失败，服务器错误", nil)
		println("// controllers/room_controller.go 获取聊天室列表失败，服务器错误 >>> err:", err.Error())
		return
	}

	respond(c, 0, "获取聊天室列表成功", rooms)
}

// AddUserToRoom 将用户添加到聊天室
func AddUserToRoom(c *gin.Context) {
	type RequestBody struct {
		UserID uint   `json:"user_id"`
		RoomID uint   `json:"room_id"`
		Role   string `json:"role"`
	}
	var body RequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		respond(c, 1, "添加用户到聊天室失败，请求错误", nil)
		println("// controllers/room_controller.go 添加用户到聊天室失败，请求错误 >>> err:", err.Error())
		return
	}

	// 检查用户是否有权限：只有在聊天室中的人才能添加用户
	userId := c.MustGet("userID").(uint)
	if !userRoomRepo.IfUserInRoom(userId, body.RoomID) {
		respond(c, 1, "添加用户到聊天室失败，没有权限，当前用户不在此房间中", nil)
		println("// controllers/room_controller.go 添加用户到聊天室失败，没有权限 >>> err:")
		return
	}

	// 检查用户是否已经在聊天室中
	if userRoomRepo.IfUserInRoom(body.UserID, body.RoomID) {
		respond(c, 1, "添加用户到聊天室失败，用户已在聊天室中", nil)
		println("// controllers/room_controller.go 添加用户到聊天室失败，用户已在聊天室中 >>> err:")
		return
	}

	// 添加用户到聊天室
	if err := userRoomRepo.AddUserToRoom(body.UserID, body.RoomID, body.Role); err != nil {
		respond(c, 1, "添加用户到聊天室失败，服务器错误", nil)
		println("// controllers/room_controller.go 添加用户到聊天室失败，服务器错误 >>> err:", err.Error())
		return
	}

	respond(c, 0, "添加用户到聊天室成功", nil)
}

// GetRoomMembers 获取聊天室成员列表
func GetRoomMembers(c *gin.Context) {
	type RequestBody struct {
		RoomID uint `json:"room_id"`
	}
	var body RequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		respond(c, 1, "获取聊天室成员列表失败，请求错误", nil)
		println("// controllers/room_controller.go 获取聊天室成员列表失败，请求错误 >>> err:", err.Error())
		return
	}

	// 检查用户是否有权限：只有在聊天室中的人才能查看成员列表
	userId := c.MustGet("userID").(uint)
	if !userRoomRepo.IfUserInRoom(userId, body.RoomID) {
		respond(c, 1, "获取聊天室成员列表失败，没有权限，当前用户不在此房间中", nil)
		println("// controllers/room_controller.go 获取聊天室成员列表失败，没有权限 >>> err:")
		return
	}

	members, err := userRoomRepo.GetRoomUsers(body.RoomID)
	if err != nil {
		respond(c, 1, "获取聊天室成员列表失败，服务器错误", nil)
		println("// controllers/room_controller.go 获取聊天室成员列表失败，服务器错误 >>> err:", err.Error())
		return
	}

	type ResponseType struct {
		ID       uint   `json:"ID"`
		Username string `json:"username"`
	}

	numbersResponse := make([]ResponseType, 0)
	for _, member := range members {
		numbersResponse = append(numbersResponse, ResponseType{
			ID:       member.ID,
			Username: member.Username,
		})
	}

	respond(c, 0, "获取聊天室成员列表成功", members)
}
