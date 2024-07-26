// routes/routes.go

package routes

import (
	"backend/internal/controllers"
	"backend/internal/middlewares"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// 跨域问题
	// 跨域中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:  []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders: []string{"Content-Type"},
		//凭证共享,确定共享
		AllowCredentials: true,
		//容许跨域的原点网站,可以直接return true 万事大吉
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		//超时时间设定
		MaxAge: 24 * time.Hour,
	}))

	// 添加 Cross-Origin-Resource-Policy 头部，解决资源跨域问题
	r.Use(func(c *gin.Context) {
		c.Header("Cross-Origin-Resource-Policy", "cross-origin")
		c.Next()
		c.Header("Cross-Origin-Resource-Policy", "cross-origin")
	})

	// 静态文件路由
	r.Static("/static", "./static")

	// Base64解码中间件
	//r.Use(middlewares.CipherDecryptMiddleware())

	// 用户注册和登录路由
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// 聊天室路由
	// 之后的操作需要用户认证，经过AuthCheck中间件
	api := r.Group("/api")
	api.Use(middlewares.AuthMiddleware())
	api.Use(middlewares.CipherDecryptMiddleware())

	// 用户信息相关路由
	user := api.Group("/user")
	{
		// 获取用户信息
		user.GET("/my", controllers.GetMyInfo)
		// 更新用户信息
		user.POST("/update", controllers.UpdateUserInfo)
		// 上传头像
		user.POST("/upload_avatar", controllers.UploadAvatar)
		// 更新密码
		user.POST("/update_password", controllers.UpdatePassword)
		// 获取指定用户信息
		user.GET("/info", controllers.GetUserInfo)
		// 获取用户好友列表
		user.GET("/friends", controllers.GetFriends)
		// 获取共同聊天室
		user.GET("/common_rooms", controllers.GetCommonRooms)
		// 获取统计信息
		user.GET("/stats", controllers.GetStats)
	}

	// 聊天室相关路由
	room := api.Group("/room")
	{
		// 创建聊天室
		room.POST("/create", controllers.CreateRoom)
		// 删除聊天室
		room.DELETE("/delete", controllers.DeleteRoom)
		// 获取聊天室信息
		room.GET("/info", controllers.GetRoomInfo)
		// 修改聊天室信息
		room.POST("/update", controllers.UpdateRoom)
		// 添加用户到聊天室
		room.POST("/add_user", controllers.AddUserToRoom)
		// 获取我的聊天室列表
		room.GET("/list", controllers.GetRoomList)
		// 获取聊天室成员列表
		room.GET("/members", controllers.GetRoomMembers)
		// 获取聊天室信息
		//room.GET("/info", controllers.GetRoomInfo)
		// 加入聊天室
		//room.POST("/join", controllers.JoinRoom)

		// 退出聊天室
		//room.POST("/exit", controllers.ExitRoom)
		// 获取聊天室消息历史（从某时间点开始，向上，第a条到第b条）
		//room.GET("/messages", controllers.GetRoomHistory)
	}

	// 聊天记录相关路由
	message := api.Group("/message")
	{
		// 发送消息
		message.POST("/send", controllers.SendMessage)
		// 获取历史消息
		message.POST("/list", controllers.GetMessages)
		// 接受上传的文件
		message.POST("/upload", controllers.UploadFile)
	}

	ws := r.Group("/ws")
	{
		// 进入WebSocket实时聊天室
		ws.GET("/join", controllers.WebSocketMessage)
	}

	return r
}
