package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"size:64;uniqueIndex;not null" json:"username"`
	Password  string         `gorm:"size:255;not null" json:"-"`
	Role      string         `gorm:"size:32;default:user" json:"role"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type UserListRequest struct {
	Page     int    `json:"page" validate:"min=1"`
	PageSize int    `json:"page_size" validate:"min=1,max=100"`
	Username string `json:"username"`
}

type UserListResponse struct {
	Total int64  `json:"total"`
	Users []User `json:"users"`
}

type CreateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role"`
}

type UpdateUserRequest struct {
	ID       uint   `json:"id" validate:"required"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
