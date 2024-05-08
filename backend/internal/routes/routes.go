// routes/routes.go

package routes

import (
	"IMTest/internal/controllers"
	"IMTest/internal/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func SetupRouter() *gin.Engine {
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

	// 用户注册和登录路由
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// 聊天室路由
	// 之后的操作需要用户认证，经过AuthCheck中间件
	api := r.Group("/api")
	api.Use(middlewares.AuthMiddleware())

	// 聊天室相关路由
	room := api.Group("/room")
	{
		// 创建聊天室
		room.POST("/create", controllers.CreateRoom)
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
	//message := api.Group("/message")
	{
		// 发送消息
		//message.POST("/send", controllers.SendMessage)
	}

	return r
}
