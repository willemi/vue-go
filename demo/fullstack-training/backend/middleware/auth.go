package middleware

import (
	"fullstack-backend/model"
	"fullstack-backend/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 验证 JWT 令牌
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, model.ErrorResponse(401, "Missing token"))
			c.Abort()
			return
		}

		// 移除 "Bearer " 前缀
		if strings.HasPrefix(tokenString, "Bearer ") {
			tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		}

		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, model.ErrorResponse(401, "Invalid or expired token"))
			c.Abort()
			return
		}

		// 将用户信息存入 context
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		c.Next()
	}
}

// AdminMiddleware 检查用户是否具有管理员角色
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