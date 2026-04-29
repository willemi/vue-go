// Package model 定义数据模型、请求/响应数据结构
package model

import (
	"time"

	"gorm.io/gorm"
)

// Menu 菜单模型，对应数据库中的 menus 表
// 支持树形结构（通过 ParentID 指向父菜单 ID）
// Role 字段控制哪些角色可以看到该菜单，用逗号分隔，如 "admin,user"
type Menu struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Title     string         `gorm:"size:64;not null" json:"title"`
	Path      string         `gorm:"size:128;uniqueIndex;not null" json:"path"`
	Icon      string         `json:"icon"`
	ParentID  uint           `json:"parent_id"`
	Sort      int            `json:"sort"`
	Hidden    bool           `json:"hidden"`
	Role      string         `json:"role"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}