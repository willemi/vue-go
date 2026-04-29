// Package database 提供数据库连接初始化和数据迁移功能
// 负责与 MySQL 建立连接、自动创建/更新数据表、初始化菜单种子数据
package database

import (
	"fullstack-backend/config"
	"fullstack-backend/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitDB 初始化数据库连接
// DSN（Data Source Name）包含数据库连接信息：
//   root:352608ww      - MySQL 用户名和密码
//   localhost:3306     - 数据库地址和端口
//   fullstack_db       - 数据库名称
//   utf8mb4            - 支持完整 Unicode（包括 emoji）
//   parseTime=True     - 自动将 MySQL 的 datetime 解析为 Go 的 time.Time
//   loc=Local          - 使用本地时区
func InitDB() error {
	dsn := "root:352608ww@tcp(localhost:3306)/fullstack_db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	config.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// AutoMigrate 根据模型定义自动创建或更新数据表结构
	// 只会添加缺失的列，不会删除或修改现有数据
	config.DB.AutoMigrate(&model.User{}, &model.Menu{})

	// 初始化菜单种子数据（仅在表为空时执行）
	seedMenus()

	return nil
}

// seedMenus 初始化菜单种子数据
// 创建两级菜单结构：Dashboard 和系统管理（一级），
// 其中"系统管理"下包含用户管理、菜单管理（二级）
func seedMenus() {
	var count int64
	config.DB.Model(&model.Menu{}).Count(&count)
	if count > 0 {
		return // 已有菜单数据，不再重复插入
	}

	// 创建一级菜单
	dashboard := model.Menu{Title: "Dashboard", Path: "/dashboard", Icon: "Odometer", ParentID: 0, Sort: 1, Hidden: false, Role: "admin,user"}
	systemMenu := model.Menu{Title: "系统管理", Path: "/system", Icon: "Setting", ParentID: 0, Sort: 2, Hidden: false, Role: "admin"}

	if err := config.DB.Create(&dashboard).Error; err != nil {
		log.Printf("Failed to seed menu '%s': %v", dashboard.Title, err)
		return
	}
	if err := config.DB.Create(&systemMenu).Error; err != nil {
		log.Printf("Failed to seed menu '%s': %v", systemMenu.Title, err)
		return
	}

	// 创建二级菜单：利用 GORM 自增 ID 特性，在 dashboard 和 systemMenu 插入后，
	// 它们的 ID 已被 GORM 填充到结构体中，可直接用作 ParentID
	userMenu := model.Menu{Title: "用户管理", Path: "/user", Icon: "User", ParentID: systemMenu.ID, Sort: 1, Hidden: false, Role: "admin,user"}
	menuMenu := model.Menu{Title: "菜单管理", Path: "/menu", Icon: "Menu", ParentID: systemMenu.ID, Sort: 2, Hidden: false, Role: "admin"}

	if err := config.DB.Create(&userMenu).Error; err != nil {
		log.Printf("Failed to seed menu '%s': %v", userMenu.Title, err)
	}
	if err := config.DB.Create(&menuMenu).Error; err != nil {
		log.Printf("Failed to seed menu '%s': %v", menuMenu.Title, err)
	}

	log.Println("Menu seed data initialized successfully")
}