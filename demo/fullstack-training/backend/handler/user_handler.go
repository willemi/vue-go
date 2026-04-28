package handler

import (
	"fullstack-backend/model"
	"fullstack-backend/service"

	"github.com/gin-gonic/gin"
)

// Login 处理用户登录
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

// GetUserList 处理获取用户列表
func GetUserList(c *gin.Context) {
	var req model.UserListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(400, model.ErrorResponse(400, "Invalid request"))
		return
	}

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

// UpdateUser 处理更新用户
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

// DeleteUser 处理删除用户
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := service.DeleteUser(id); err != nil {
		c.JSON(500, model.ErrorResponse(500, "Failed to delete user"))
		return
	}

	c.JSON(204, nil)
}
