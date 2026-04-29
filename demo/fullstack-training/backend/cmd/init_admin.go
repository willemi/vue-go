// Package main 提供初始化管理员账户的工具
// 运行此程序可在数据库中创建默认管理员账户
// 使用方式: go run cmd/init_admin.go
package main

import (
	"fmt"
	"fullstack-backend/config"
	"fullstack-backend/database"
	"fullstack-backend/model"
	"fullstack-backend/utils"
)

func main() {
	// 初始化数据库连接
	if err := database.InitDB(); err != nil {
		panic(err)
	}

	// 使用 bcrypt 生成密码哈希
	hash, err := utils.HashPassword("admin123")
	if err != nil {
		panic(err)
	}
	fmt.Println("Password hash:", hash)

	// 构建管理员用户对象
	admin := model.User{
		Username: "admin",
		Password: string(hash),
		Role:     "admin",
	}

	// 插入数据库并输出结果
	if err := config.DB.Create(&admin).Error; err != nil {
		fmt.Println("Insert failed:", err)
	} else {
		fmt.Println("Admin user created successfully! ID:", admin.ID)
	}
}