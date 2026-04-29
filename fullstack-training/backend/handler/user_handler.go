// Package handler 包含所有 HTTP 请求处理函数（Controller 层）
// 每个函数处理一种 API 请求：解析参数 -> 调用 service 层 -> 返回 JSON 响应
package handler

import (
	"fullstack-backend/model"
	"fullstack-backend/service"

	"github.com/gin-gonic/gin"
)

// Login 处理用户登录
// POST /api/user/login
// 请求体：{ username, password }
// 成功响应：{ code: 200, data: { token, user } }
// 失败响应：{ code: 401, message: "Invalid username or password" }
func Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, model.ErrorResponse(400, "Invalid request"))
		return
	}

	token, user, err := service.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(401, model.ErrorResponse(401, "Invalid username or password"))
		return
	}

	c.JSON(200, model.SuccessResponse(model.LoginResponse{
		Token: token,
		User:  user,
	}))
}

// GetUserList 处理获取用户列表（分页）
// GET /api/user/list?username=&page=1&page_size=10
// 支持按用户名模糊搜索，默认 page=1, page_size=10
// 响应：{ code: 200, data: { total, users } }
func GetUserList(c *gin.Context) {
	var req model.UserListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(400, model.ErrorResponse(400, "Invalid request"))
		return
	}

	// 设置分页默认值
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}

	users, total, err := service.GetUserList(req.Username, req.Page, req.PageSize)
	if err != nil {
		c.JSON(500, model.ErrorResponse(500, "Failed to get user list"))
		return
	}

	c.JSON(200, model.SuccessResponse(model.UserListResponse{
		Total: total,
		Users: users,
	}))
}

// CreateUser 处理创建新用户
// POST /api/user/add
// 请求体：{ username, password, role }
// 成功响应：{ code: 201, data: user }
// 失败响应：{ code: 400/500, message: "..." }
func CreateUser(c *gin.Context) {
	var req model.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, model.ErrorResponse(400, "Invalid request"))
		return
	}

	if req.Username == "" || req.Password == "" {
		c.JSON(400, model.ErrorResponse(400, "Username and password are required"))
		return
	}

	user, err := service.CreateUser(req)
	if err != nil {
		c.JSON(500, model.ErrorResponse(500, "Failed to create user"))
		return
	}

	c.JSON(201, model.SuccessResponse(user))
}

// UpdateUser 处理更新用户信息
// PUT /api/user/edit
// 请求体：{ id, username, password?, role }
// password 为可选字段，空值表示不修改密码
// 成功响应：{ code: 200, data: user }
func UpdateUser(c *gin.Context) {
	var req model.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, model.ErrorResponse(400, "Invalid request"))
		return
	}

	if req.ID == 0 {
		c.JSON(400, model.ErrorResponse(400, "User ID is required"))
		return
	}

	user, err := service.UpdateUser(req)
	if err != nil {
		c.JSON(500, model.ErrorResponse(500, "Failed to update user"))
		return
	}

	c.JSON(200, model.SuccessResponse(user))
}

// DeleteUser 处理删除用户（软删除）
// DELETE /api/user/delete/:id
// GORM 软删除：将 DeletedAt 字段设为当前时间，而非真正从数据库删除
// 成功响应：{ code: 204, data: null }
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := service.DeleteUser(id); err != nil {
		c.JSON(500, model.ErrorResponse(500, "Failed to delete user"))
		return
	}

	c.JSON(204, nil)
}