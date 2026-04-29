// Package main 后台管理系统的程序入口
// 负责初始化配置、启动数据库连接和 HTTP 服务器
package main

import (
	"fullstack-backend/config"
	"fullstack-backend/server"
)

// main 程序主入口
// 执行流程：
// 1. 调用 config.Init() 加载全局配置（当前仅端口配置）
// 2. 调用 server.New() 创建 Gin 引擎并初始化数据库
// 3. 调用 r.Run(":8080") 启动 HTTP 服务器监听 8080 端口
func main() {
	// 初始化配置
	config.Init()

	// 初始化服务器
	r := server.New()

	// 启动服务器
	r.Run(":8080")
}