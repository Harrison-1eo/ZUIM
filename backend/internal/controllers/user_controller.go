// controllers/user_controller.go

package controllers

import (
	"backend/internal/models"
	"backend/internal/repositories"
	"backend/internal/utils"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var UserEncryptKey = map[uint][4]string{}
var UserCipherBackends = map[uint]*utils.StreamCipher{}
var UserCipherFrontends = map[uint]*utils.StreamCipher{}
var UserCipherWebsocketsBackends = map[uint]*utils.StreamCipher{}
var UserCipherWebsocketsFrontends = map[uint]*utils.StreamCipher{}

var keySalts = []string{"backend", "frontend", "websocketBackend", "websocketFrontend"}

var userRepository = repositories.NewUserRepository()
var statsRepository = repositories.NewStatsRepository()

// Register 注册用户
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		respondWithoutEncrypt(c, 1, "注册失败，请求错误", nil)
		println("// controllers/user_controller.go 注册失败，请求错误 >>> err:", err.Error())
		return
	}

	// 检查用户名是否已存在
	if userRepository.UserExists(user.Username) {
		respondWithoutEncrypt(c, 1, "注册失败，用户名已存在", nil)
		println("// controllers/user_controller.go 注册失败，用户名已存在 >>> err:")
		return
	}

	// 创建用户
	newUser, err := userRepository.CreateUser(user)
	if err != nil {
		respondWithoutEncrypt(c, 1, "注册失败，服务器错误", nil)
		println("// controllers/user_controller.go 注册失败，服务器错误 >>> err:", err.Error())
		return
	}

	_, err = userInfoRepository.CreateUserInfo(models.UserInfo{
		UserID:   newUser.ID,
		Username: newUser.Username,
		Avatar:   "/static/avatars/nopic.png",
	})

	if err != nil {
		respondWithoutEncrypt(c, 1, "注册失败，服务器错误", nil)
		println("// controllers/user_controller.go 注册失败，服务器错误 >>> err:", err.Error())
		return
	}

	type ResponseType struct {
		ID       uint   `json:"ID"`
		Username string `json:"username"`
	}

	respondWithoutEncrypt(c, 0, "注册成功",
		ResponseType{ID: newUser.ID,
			Username: newUser.Username})
}

// Login 用户登录
func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		respondWithoutEncrypt(c, 1, "登录失败，请求错误", nil)
		println("// controllers/user_controller.go 登录失败，请求错误 >>> err:", err.Error())
		return
	}

	// 验证用户登录信息
	authenticatedUser, err := userRepository.AuthenticateUser(user.Username, user.Password)
	if err != nil {
		respondWithoutEncrypt(c, 1, "登录失败，用户名或密码错误", nil)
		println("// controllers/user_controller.go 登录失败，用户名或密码错误 >>> err:", err.Error())
		return
	}

	// 检查用户信息是否存在，若不存在则创建
	if _, err := userInfoRepository.GetUserInfoByID(authenticatedUser.ID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			_, err := userInfoRepository.CreateUserInfo(models.UserInfo{
				UserID:   authenticatedUser.ID,
				Username: authenticatedUser.Username,
			})
			if err != nil {
				respondWithoutEncrypt(c, 1, "登录失败，服务器错误", nil)
				println("// controllers/user_controller.go 登录失败，服务器错误 >>> err:", err.Error())
				return
			}
		} else {
			respondWithoutEncrypt(c, 1, "登录失败，服务器错误", nil)
			println("// controllers/user_controller.go 登录失败，服务器错误 >>> err:", err.Error())
			return
		}
	}

	// 生成JWT令牌
	token, err := utils.CreateToken(authenticatedUser.ID)
	if err != nil {
		respondWithoutEncrypt(c, 1, "登录失败，服务器错误", nil)
		println("// controllers/user_controller.go 登录失败，服务器错误 >>> err:", err.Error())
		return
	}

	println("create token: ", token)

	var userKeys [4]string
	for index, salt := range keySalts {
		k := token + utils.Sha256OfString(authenticatedUser.Password) + strconv.Itoa(int(authenticatedUser.ID)) + salt
		userKeys[index] = (utils.Sha256OfString(k))
	}
	UserEncryptKey[authenticatedUser.ID] = userKeys

	UserCipherBackends[authenticatedUser.ID] = utils.NewStreamCipher(userKeys[0])
	UserCipherFrontends[authenticatedUser.ID] = utils.NewStreamCipher(userKeys[1])
	UserCipherWebsocketsBackends[authenticatedUser.ID] = utils.NewStreamCipher(userKeys[2])
	UserCipherWebsocketsFrontends[authenticatedUser.ID] = utils.NewStreamCipher(userKeys[3])

	println(" ++++ UserEncryptKey ++++ ")
	println("UserEncryptKey[authenticatedUser.ID]: ", authenticatedUser.ID)
	println(UserEncryptKey[authenticatedUser.ID][0], UserCipherBackends[authenticatedUser.ID].Key)
	println(UserEncryptKey[authenticatedUser.ID][1], UserCipherFrontends[authenticatedUser.ID].Key)
	println(UserEncryptKey[authenticatedUser.ID][2], UserCipherWebsocketsBackends[authenticatedUser.ID].Key)
	println(UserEncryptKey[authenticatedUser.ID][3], UserCipherWebsocketsFrontends[authenticatedUser.ID].Key)
	println(" ++++ UserEncryptKey ++++ \n")

	type ResponseType struct {
		ID       uint   `json:"ID"`
		Username string `json:"username"`
	}

	logIn := ResponseType{ID: authenticatedUser.ID,
		Username: authenticatedUser.Username}

	respondWithoutEncrypt(c, 0, "登录成功", gin.H{"user": logIn, "token": token})
}

func UpdatePassword(c *gin.Context) {
	type RequestBody struct {
		Username    string `json:"username"`
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}
	var reqBody RequestBody
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		respond(c, 1, "修改密码失败，请求错误", nil)
		println("// controllers/user_controller.go 修改密码失败，请求错误 >>> err:", err.Error())
		return
	}

	// 验证用户登录信息
	authenticatedUser, err := userRepository.AuthenticateUser(reqBody.Username, reqBody.OldPassword)
	if err != nil {
		respond(c, 1, "修改密码失败，原密码错误", nil)
		println("// controllers/user_controller.go 修改密码失败，用户名或密码错误 >>> err:", err.Error())
		return
	}

	// 更新密码
	err = userRepository.UpdatePassword(authenticatedUser.ID, reqBody.NewPassword)
	if err != nil {
		respond(c, 1, "修改密码失败，服务器错误", nil)
		println("// controllers/user_controller.go 修改密码失败，服务器错误 >>> err:", err.Error())
		return
	}

	respond(c, 0, "修改密码成功", nil)
}

func GetStats(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	// 查询用户在多少个房间中
	roomCount, err := statsRepository.GetUserRoomCount(userID)
	if err != nil {
		respond(c, 1, "获取统计信息（用户所在房间数）失败", nil)
		println("// controllers/user_controller.go 获取统计信息（用户所在房间数）失败 >>> err:", err.Error())
		return
	}

	// 查询用户有多少个好友
	friendCount, err := statsRepository.GetUserFriendCount(userID)
	if err != nil {
		respond(c, 1, "获取统计信息（用户好友数）失败", nil)
		println("// controllers/user_controller.go 获取统计信息（用户好友数）失败 >>> err:", err.Error())
		return
	}

	// 查询在线用户
	onlineUserCount, err := statsRepository.GetOnlineUserCount()
	if err != nil {
		respond(c, 1, "获取统计信息（在线用户数）失败", nil)
		println("// controllers/user_controller.go 获取统计信息（在线用户数）失败 >>> err:", err.Error())
		return
	}

	// 查看未读消息数
	unreadMessageCount, err := statsRepository.GetUnreadMessageCount(userID)
	if err != nil {
		respond(c, 1, "获取统计信息（未读消息数）失败", nil)
		println("// controllers/user_controller.go 获取统计信息（未读消息数）失败 >>> err:", err.Error())
		return
	}

	// 查询最近一周的该用户每天发送的消息数量
	messageCount, err := statsRepository.GetUserMessageCount(userID)
	if err != nil {
		respond(c, 1, "获取统计信息（用户消息数）失败", nil)
		println("// controllers/user_controller.go 获取统计信息（用户消息数）失败 >>> err:", err.Error())
		return
	}

	// 查询该用户在的各个房间中的所有消息数量分别是多少
	roomMessageCount, err := statsRepository.GetUserAllMessageCountInRooms(userID)
	if err != nil {
		respond(c, 1, "获取统计信息（房间消息数）失败", nil)
		println("// controllers/user_controller.go 获取统计信息（房间消息数）失败 >>> err:", err.Error())
		return
	}

	roomNameMessageCount := map[string]int{}
	for roomID, count := range roomMessageCount {
		room, err := roomRepo.GetRoomWithDeleted(roomID)
		if err != nil {
			respond(c, 1, "获取统计信息（房间消息数）失败1", nil)
			println("// controllers/user_controller.go 获取统计信息（房间消息数）失败1 >>> err:", err.Error())
			return
		}
		roomNameMessageCount[room.Name] = count
	}

	type ResponseType struct {
		RoomCount          int            `json:"room_count"`
		FriendCount        int            `json:"friend_count"`
		OnlineUserCount    int            `json:"online_user_count"`
		UnreadMessageCount int            `json:"unread_message_count"`
		MessageCount       map[string]int `json:"message_count"`
		RoomMessageCount   map[string]int `json:"room_message_count"`
	}

	respond(c, 0, "获取统计信息成功", ResponseType{
		RoomCount:          roomCount,
		FriendCount:        friendCount,
		OnlineUserCount:    onlineUserCount,
		UnreadMessageCount: unreadMessageCount,
		MessageCount:       messageCount,
		RoomMessageCount:   roomNameMessageCount,
	})

}
