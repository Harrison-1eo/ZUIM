// cmd/main.go

package main

import (
	"backend/internal/repositories"
	"backend/internal/routes"
)

func main() {
	// 连接SQLite数据库
	//repositories.InitDBSqlite()

	// 连接MySQL数据库
	repositories.InitDBMySQL()

	router := routes.SetupRouter()
	err := router.Run(":8000")
	if err != nil {
		return
	}
}
