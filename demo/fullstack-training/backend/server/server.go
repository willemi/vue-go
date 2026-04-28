package server

import (
	"fullstack-backend/database"
	"fullstack-backend/handler"
	"fullstack-backend/middleware"

	"github.com/gin-gonic/gin"
)

// New 创建并配置所有路由的 Gin 服务器
func New() *gin.Engine {
	// 初始化数据库
	if err := database.InitDB(); err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	r := gin.Default()

	// 应用 CORS 中间件
	r.Use(middleware.CORSMiddleware())

	// 公开路由（无需认证）
	r.POST("/api/user/login", handler.Login)

	// 受保护路由（需要认证）
	auth := r.Group("/api")
	auth.Use(middleware.AuthMiddleware())
	{
		// 用户管理路由
		auth.GET("/user/list", handler.GetUserList)
		auth.POST("/user/add", handler.CreateUser)
		auth.PUT("/user/edit", handler.UpdateUser)
		auth.DELETE("/user/delete/:id", handler.DeleteUser)

		// 菜单管理路由
		auth.GET("/menu/list", handler.GetMenuList)
		auth.GET("/menu/tree", handler.GetMenuTree)
		auth.POST("/menu/add", handler.CreateMenu)
		auth.PUT("/menu/edit", handler.UpdateMenu)
		auth.DELETE("/menu/delete/:id", handler.DeleteMenu)
	}

	// 仅管理员路由
	admin := r.Group("/api")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		admin.DELETE("/admin/user/delete/:id", handler.DeleteUser)
	}

	return r
}
