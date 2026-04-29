// Package utils 提供加密、签名等工具函数
package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword 使用 bcrypt 对明文密码进行哈希
// bcrypt.DefaultCost 自动选择合适的计算成本（当前为 10）
// 返回哈希后的字符串和可能的错误
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPassword 验证明文密码与哈希值是否匹配
// 通过 bcrypt.CompareHashAndPassword 比较，匹配返回 true，不匹配返回 false
// 避免了时序攻击（timing attack）
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}