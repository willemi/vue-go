// Package middleware 提供 Gin HTTP 中间件
// 中间件是请求处理链中的拦截器，用于在调用 handler 之前执行认证、权限检查等逻辑
package middleware

import (
	"fullstack-backend/model"
	"fullstack-backend/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware JWT 认证中间件
// 检查请求头中是否包含有效的 Authorization: Bearer <token>
// 验证通过后，将用户信息（user_id、username、role）存入 Gin Context
// 后续 handler 可通过 c.Get("xxx") 获取这些信息
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, model.ErrorResponse(401, "Missing token"))
			c.Abort()
			return
		}

		// 移除 "Bearer " 前缀（不区分大小写），保留 token 本身
		if strings.HasPrefix(tokenString, "Bearer ") {
			tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		}

		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, model.ErrorResponse(401, "Invalid or expired token"))
			c.Abort()
			return
		}

		// 将解析出的用户信息存入 Context，供后续 handler 使用
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		// 调用 c.Next() 继续执行后续中间件和 handler
		c.Next()
	}
}

// AdminMiddleware 管理员权限检查中间件
// 应在 AuthMiddleware 之后使用，因为依赖 AuthMiddleware 设置的 role 值
// 仅当 role 为 "admin" 时允许通过，否则返回 403 Forbidden
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != "admin" {
			c.JSON(http.StatusForbidden, model.ErrorResponse(403, "Permission denied"))
			c.Abort()
			return
		}
		c.Next()
	}
}