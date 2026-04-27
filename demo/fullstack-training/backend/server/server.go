package server

import (
	"fullstack-backend/database"
	"fullstack-backend/handler"
	"fullstack-backend/middleware"

	"github.com/gin-gonic/gin"
)

// New creates a new Gin server with all routes configured
func New() *gin.Engine {
	// Initialize database
	if err := database.InitDB(); err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	r := gin.Default()

	// Apply CORS middleware
	r.Use(middleware.CORSMiddleware())

	// Public routes (no auth required)
	r.POST("/api/user/login", handler.Login)

	// Protected routes (auth required)
	auth := r.Group("/api")
	auth.Use(middleware.AuthMiddleware())
	{
		// User management routes
		auth.GET("/user/list", handler.GetUserList)
		auth.POST("/user/add", handler.CreateUser)
		auth.PUT("/user/edit", handler.UpdateUser)
		auth.DELETE("/user/delete/:id", handler.DeleteUser)

		// Menu management routes
		auth.GET("/menu/list", handler.GetMenuList)
		auth.POST("/menu/add", handler.CreateMenu)
		auth.PUT("/menu/edit", handler.UpdateMenu)
		auth.DELETE("/menu/delete/:id", handler.DeleteMenu)
	}

	// Admin-only routes
	admin := r.Group("/api")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		admin.DELETE("/admin/user/delete/:id", handler.DeleteUser)
	}

	return r
}