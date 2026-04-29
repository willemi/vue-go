// Package middleware 提供 Gin HTTP 中间件
package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORSMiddleware 处理跨域资源共享（CORS）
// 允许前端开发服务器（localhost:5173/3000）向 API 发起跨域请求
//
// 配置说明：
//   - AllowOrigins: 允许的来源域名，开发环境为 Vite 和部分常见端口
//   - AllowMethods: 允许的 HTTP 方法
//   - AllowHeaders: 允许的请求头，Authorization 是发送 JWT token 必需的
//   - AllowCredentials: 允许携带 Cookie（注意：此时 AllowOrigins 不能为 *）
//   - MaxAge: 预检请求（OPTIONS）的缓存时间，减少 OPTIONS 请求次数
func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}