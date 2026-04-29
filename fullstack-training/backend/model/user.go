// Package model 定义数据模型、请求/响应数据结构
// 包含 User、Menu 两个业务模型，以及通用的 API 响应包装
package model

import (
	"time"

	"gorm.io/gorm"
)

// User 用户模型，对应数据库中的 users 表
// - Username: 用户名，唯一索引，不能为空
// - Password: 密码，经过 bcrypt 哈希后存储，JSON 序列化时隐藏
// - Role: 角色，可选 admin 或 user，默认为 user
// - DeletedAt: GORM 软删除字段，删除时标记时间而非物理删除
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"size:64;uniqueIndex;not null" json:"username"`
	Password  string         `gorm:"size:255;not null" json:"-"`
	Role      string         `gorm:"size:32;default:user" json:"role"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// LoginRequest 登录请求结构
// 前端提交 username 和 password 进行身份验证
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// LoginResponse 登录成功响应结构
// 包含 JWT 令牌和用户基本信息
type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

// UserListRequest 获取用户列表的请求参数
// 支持按用户名模糊搜索，并分页返回结果
type UserListRequest struct {
	Page     int    `json:"page" validate:"min=1"`
	PageSize int    `json:"page_size" validate:"min=1,max=100"`
	Username string `json:"username"`
}

// UserListResponse 用户列表响应结构
// Total: 符合条件的用户总数，Users: 当前页的用户列表
type UserListResponse struct {
	Total int64  `json:"total"`
	Users []User `json:"users"`
}

// CreateUserRequest 创建新用户的请求结构
// Username 和 Password 为必填项，Role 默认为 user
type CreateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role"`
}

// UpdateUserRequest 更新用户的请求结构
// ID 为必填，其他字段可选（只更新非空字段）
type UpdateUserRequest struct {
	ID       uint   `json:"id" validate:"required"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}