package model

import (
	"time"

	"gorm.io/gorm"
)

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