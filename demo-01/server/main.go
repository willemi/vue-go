package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 允许所有跨域（开发用）
	// 开发环境：Go 端加一行 cors.Default() 即可
	// 生产环境：用 Nginx 把前端和后端代理到同一个域名，完全无跨域
	r.Use(cors.Default())

	// 测试接口
	r.GET("/api/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "来自 Golang 的问候",
			"data":    nil,
		})
	})

	// 用户信息接口
	r.GET("/api/user/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": gin.H{
				"username": "vue-go-user",
				"nickname": "测试用户",
				"age":      20,
			},
		})
	})

	// POST 示例
	r.POST("/api/login", func(c *gin.Context) {
		var req struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "参数错误"})
			return
		}

		if req.Username == "admin" && req.Password == "123456" {
			c.JSON(http.StatusOK, gin.H{
				"code":  0,
				"token": "jwt-token-abc123",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "账号或密码错误"})
		}
	})

	// 启动在 8080
	r.Run(":8080")
}
