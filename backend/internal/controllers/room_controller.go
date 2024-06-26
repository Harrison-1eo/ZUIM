// controllers/room_controller.go

package controllers

import (
	"backend/internal/models"
	"backend/internal/repositories"
	"github.com/gin-gonic/gin"
	"strconv"
)

var roomRepo *repositories.RoomRepository = repositories.NewRoomRepository()
var userRoomRepo *repositories.UserRoomRepository = repositories.NewUserRoomRepository()
var userRepo *repositories.UserRepository = repositories.NewUserRepository()

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

// GetRoomInfo 获取聊天室信息
func GetRoomInfo(c *gin.Context) {
	// 从URL参数中解析出聊天室ID
	roomID, ok := c.GetQuery("room_id")
	if ok != true {
		respond(c, 1, "获取聊天室信息失败，请求错误", nil)
		println("// controllers/room_controller.go 获取聊天室信息失败，请求错误 >>> err: room_id not found", c.Request.URL.RawQuery)
		return
	}
	roomIDUint, err := strconv.Atoi(roomID)
	if roomIDUint == 0 || err != nil {
		respond(c, 1, "获取聊天室信息失败，请求错误", nil)
		println("// controllers/room_controller.go 获取聊天室信息失败，请求错误 >>> err: roomID == 0")
		return
	}

	// 检查用户是否有权限：只有在聊天室中的人才能查看聊天室信息
	userId := c.MustGet("userID").(uint)
	if !userRoomRepo.IfUserInRoom(userId, uint(roomIDUint)) {
		respond(c, 1, "获取聊天室信息失败，没有权限，当前用户不在此房间中", nil)
		println("// controllers/room_controller.go 获取聊天室信息失败，没有权限 >>> err:")
		return
	}

	room, err := roomRepo.GetRoom(uint(roomIDUint))
	if err != nil {
		respond(c, 1, "获取聊天室信息失败，服务器错误", nil)
		println("// controllers/room_controller.go 获取聊天室信息失败，服务器错误 >>> err:", err.Error())
		return
	}

	respond(c, 0, "获取聊天室信息成功", room)
}

// UpdateRoom 修改聊天室信息
func UpdateRoom(c *gin.Context) {
	type RequestBody struct {
		RoomID   uint   `json:"room_id"`
		RoomName string `json:"name"`
		RoomInfo string `json:"description"`
	}

	// 从请求体中解析出聊天室信息
	var body RequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		respond(c, 1, "修改聊天室信息失败，请求错误", nil)
		println("// controllers/room_controller.go 修改聊天室信息失败，请求错误 >>> err:", err.Error())
		return
	}

	// 检查用户是否有权限：只有聊天室中的用户才能修改聊天室信息
	userId := c.MustGet("userID").(uint)
	if !userRoomRepo.IfUserInRoom(userId, body.RoomID) {
		respond(c, 1, "修改聊天室信息失败，没有权限，当前用户不在此房间中", nil)
		println("// controllers/room_controller.go 修改聊天室信息失败，没有权限 >>> err:")
		return
	}

	// 获取聊天室信息
	room, err := roomRepo.GetRoom(body.RoomID)
	if err != nil {
		respond(c, 1, "修改聊天室信息失败，聊天室不存在", nil)
		println("// controllers/room_controller.go 修改聊天室信息失败，服务器错误1 >>> err:", err.Error())
		return
	}

	// 修改聊天室信息
	room.Name = body.RoomName
	room.Description = body.RoomInfo

	newRoom, err := roomRepo.UpdateRoom(*room)
	if err != nil {
		respond(c, 1, "修改聊天室信息失败，服务器错误2", nil)
		println("// controllers/room_controller.go 修改聊天室信息失败，服务器错误2 >>> err:", err.Error())
		return
	}

	respond(c, 0, "修改聊天室信息成功", newRoom)
}

// DeleteRoom 根据RoomID删除聊天室
func DeleteRoom(c *gin.Context) {
	// 从URL参数中解析出聊天室ID
	roomID, ok := c.GetQuery("room_id")

	if ok != true {
		respond(c, 1, "删除聊天室失败，请求错误1", nil)
		println("// controllers/room_controller.go 删除聊天室失败，请求错误1 >>> err: room_id not found", c.Request.URL.RawQuery)
		//println(roomID)
		return
	}
	roomIDUint, err := strconv.Atoi(roomID)
	if roomIDUint == 0 || err != nil {
		respond(c, 1, "删除聊天室失败，请求错误2", nil)
		println("// controllers/room_controller.go 删除聊天室失败，请求错误2 >>> err: roomID == 0")
		return
	}

	// // 检查用户是否有权限：只有聊天室的创建者才能删除聊天室
	// userId := c.MustGet("userID").(uint)
	// if !userRoomRepo.IfUserIsOwner(userId, uint(roomIDUint)) {
	// 	respond(c, 1, "删除聊天室失败，没有权限，当前用户不是此房间的创建者", nil)
	// 	println("// controllers/room_controller.go 删除聊天室失败，没有权限 >>> err:")
	// 	return
	// }

	// 删除聊天室
	if err := roomRepo.DeleteRoom(uint(roomIDUint)); err != nil {
		//println(roomID)
		respond(c, 1, "删除聊天室失败，服务器错误3", nil)
		println("// controllers/room_controller.go 删除聊天室失败，服务器错误3 >>> err:", err.Error())
		return
	}
	// user_room表中删除相关记录
	if err := userRoomRepo.DeleteRoom(uint(roomIDUint)); err != nil {
		respond(c, 1, "删除聊天室失败，服务器错误4", nil)
		println("// controllers/room_controller.go 删除聊天室失败，服务器错误4 >>> err:", err.Error())
		return
	}

	respond(c, 0, "删除聊天室成功", nil)
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
		UserName string `json:"user_name"`
		RoomID   uint   `json:"room_id"`
		Role     string `json:"role"`
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

	// 检查用户是否存在
	addUser, err := userRepo.GetUserByUsername(body.UserName)
	if err != nil {
		respond(c, 1, "添加用户到聊天室失败，用户不存在", nil)
		println("// controllers/room_controller.go 添加用户到聊天室失败，用户不存在 >>> err:")
		return
	}

	// 检查用户是否已经在聊天室中
	if userRoomRepo.IfUserInRoom(addUser.ID, body.RoomID) {
		respond(c, 1, "添加用户到聊天室失败，用户已在聊天室中", nil)
		println("// controllers/room_controller.go 添加用户到聊天室失败，用户已在聊天室中 >>> err:")
		return
	}

	// 添加用户到聊天室
	if err := userRoomRepo.AddUserToRoom(addUser.ID, body.RoomID, body.Role); err != nil {
		respond(c, 1, "添加用户到聊天室失败，服务器错误", nil)
		println("// controllers/room_controller.go 添加用户到聊天室失败，服务器错误 >>> err:", err.Error())
		return
	}

	respond(c, 0, "添加用户到聊天室成功", nil)
}

// GetRoomMembers 获取聊天室成员列表
func GetRoomMembers(c *gin.Context) {
	// 从URL参数中解析出聊天室ID
	roomID, ok := c.GetQuery("room_id")
	if ok != true {
		respond(c, 1, "获取聊天室成员列表失败，请求错误", nil)
		println("// controllers/room_controller.go 获取聊天室成员列表失败，请求错误 >>> err: room_id not found", c.Request.URL.RawQuery)
		return
	}
	roomIDUint, err := strconv.Atoi(roomID)
	if roomIDUint == 0 || err != nil {
		respond(c, 1, "获取聊天室成员列表失败，请求错误", nil)
		println("// controllers/room_controller.go 获取聊天室成员列表失败，请求错误 >>> err: roomID == 0")
		return
	}

	// 检查用户是否有权限：只有在聊天室中的人才能查看成员列表
	userId := c.MustGet("userID").(uint)
	if !userRoomRepo.IfUserInRoom(userId, uint(roomIDUint)) {
		respond(c, 1, "获取聊天室成员列表失败，没有权限，当前用户不在此房间中", nil)
		println("// controllers/room_controller.go 获取聊天室成员列表失败，没有权限 >>> err:")
		return
	}

	members, err := userRoomRepo.GetRoomUserInfos(uint(roomIDUint))
	if err != nil {
		respond(c, 1, "获取聊天室成员列表失败，服务器错误", nil)
		println("// controllers/room_controller.go 获取聊天室成员列表失败，服务器错误 >>> err:", err.Error())
		return
	}

	type ResponseType struct {
		ID       uint   `json:"ID"`
		Username string `json:"username"`
		Avatar   string `json:"avatar"`
		NikeName string `json:"nike_name"`
	}

	numbersResponse := make([]ResponseType, 0)
	for _, member := range members {
		numbersResponse = append(numbersResponse, ResponseType{
			ID:       member.UserID,
			Username: member.Username,
			Avatar:   member.Avatar,
			NikeName: member.NikeName,
		})
	}

	respond(c, 0, "获取聊天室成员列表成功", numbersResponse)
}
