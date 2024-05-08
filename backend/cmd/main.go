// cmd/main.go

package main

import (
	"IMTest/internal/repositories"
	"IMTest/internal/routes"
)

func main() {
	// 连接SQLite数据库
	repositories.InitDB()

	router := routes.SetupRouter()
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
