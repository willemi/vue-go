// Package config 提供全局配置管理
// 当前主要管理数据库连接实例和 HTTP 服务端口
package config

import (
	"os"

	"gorm.io/gorm"
)

// DB 是全局数据库连接实例，供整个项目共享使用
// 在 database.InitDB() 初始化后可用
var (
	DB  *gorm.DB
	Port string
)

// Init 初始化全局配置
// 目前仅加载 HTTP 服务端口，可扩展为从配置文件或环境变量加载更多配置项
func Init() {
	Port = getEnv("PORT", "8080")
}

// getEnv 读取环境变量，若未设置则返回默认值
// key: 环境变量名称
// defaultValue: 未设置时的默认值
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}