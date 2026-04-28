package database

import (
	"fullstack-backend/config"
	"fullstack-backend/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitDB 初始化数据库连接
func InitDB() error {
	dsn := "root:352608ww@tcp(localhost:3306)/fullstack_db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	config.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// 自动迁移数据库结构
	config.DB.AutoMigrate(&model.User{}, &model.Menu{})

	return nil
}
