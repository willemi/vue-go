// Package main 提供密码哈希生成工具
// 用于在命令行中手动生成任意密码的 bcrypt 哈希值，验证密码是否正确。
// 使用方式: go run cmd/password.go
package main

import (
	"fmt"
	"fullstack-backend/utils"
)

func main() {
	// 默认生成 "admin123" 的哈希，可根据需要修改密码
	hash, err := utils.HashPassword("admin123")
	if err != nil {
		fmt.Println("Hash failed:", err)
		return
	}
	fmt.Println("Password hash:", hash)
}