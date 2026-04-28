package database

import (
	"fullstack-backend/config"
	"fullstack-backend/model"
	"log"

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

	// 初始化菜单种子数据
	seedMenus()

	return nil
}

// seedMenus 初始化菜单种子数据
func seedMenus() {
	var count int64
	config.DB.Model(&model.Menu{}).Count(&count)
	if count > 0 {
		return // 已有菜单数据，不再重复插入
	}

	// 一级菜单
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

	// 二级菜单：使用父菜单的自增 ID
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
