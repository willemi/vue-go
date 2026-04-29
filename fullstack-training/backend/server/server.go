// Package server 提供 Gin HTTP 服务器的创建与路由配置
// 包含公开路由、受保护路由和管理员专属路由三层
package server

import (
	"fullstack-backend/database"
	"fullstack-backend/handler"
	"fullstack-backend/middleware"

	"github.com/gin-gonic/gin"
)

// New 创建并配置所有路由的 Gin 服务器
// 执行顺序：连接数据库 -> 创建 Gin 实例 -> 注册中间件 -> 注册路由
// 数据库初始化失败时 panic，整个服务无法启动
func New() *gin.Engine {
	// 初始化数据库连接（GORM 自动迁移模型结构）
	if err := database.InitDB(); err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	// 创建默认 Gin 实例（内置 Logger 和 Recovery 中间件）
	r := gin.Default()

	// 应用 CORS 中间件，允许前端开发服务器跨域请求
	r.Use(middleware.CORSMiddleware())

	// ========== 公开路由（无需认证）==========

	// 用户登录接口，路径为 /api/user/login
	r.POST("/api/user/login", handler.Login)

	// ========== 受保护路由（需要 JWT 认证）==========

	// 所有 /api/* 路由组统一应用认证中间件
	auth := r.Group("/api")
	auth.Use(middleware.AuthMiddleware())
	{
		// 用户管理 CRUD
		auth.GET("/user/list", handler.GetUserList)      // 获取用户列表（分页 + 搜索）
		auth.POST("/user/add", handler.CreateUser)       // 创建用户
		auth.PUT("/user/edit", handler.UpdateUser)       // 更新用户
		auth.DELETE("/user/delete/:id", handler.DeleteUser) // 删除用户（软删除）

		// 菜单管理 CRUD
		auth.GET("/menu/list", handler.GetMenuList)      // 获取菜单列表（扁平结构）
		auth.GET("/menu/tree", handler.GetMenuTree)      // 获取当前用户可见的菜单树
		auth.POST("/menu/add", handler.CreateMenu)       // 创建菜单
		auth.PUT("/menu/edit", handler.UpdateMenu)       // 更新菜单
		auth.DELETE("/menu/delete/:id", handler.DeleteMenu) // 删除菜单（软删除）
	}

	// ========== 仅管理员路由（需要认证 + admin 角色）==========

	admin := r.Group("/api")
	// 串联 AuthMiddleware 和 AdminMiddleware：先验证身份，再检查角色
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		// 管理员专属删除操作（普通用户也可访问 /api/user/delete/:id，但管理员有独立路径）
		admin.DELETE("/admin/user/delete/:id", handler.DeleteUser)
	}

	return r
}