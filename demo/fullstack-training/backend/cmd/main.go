package main

import (
	"fullstack-backend/config"
	"fullstack-backend/server"
)

func main() {
	// 初始化配置
	config.Init()

	// 初始化服务器
	r := server.New()

	// 启动服务器
	r.Run(":8080")
}