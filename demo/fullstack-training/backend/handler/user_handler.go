package handler

import (
	"fullstack-backend/model"
	"fullstack-backend/service"

	"github.com/gin-gonic/gin"
)

// Login handles user login
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

// GetUserList handles getting user list
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

// CreateUser handles creating a new user
func CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, model.ErrorResponse(400, "Invalid request"))
		return
	}

	if err := service.CreateUser(&user); err != nil {
		c.JSON(500, model.ErrorResponse(500, "Failed to create user"))
		return
	}

	c.JSON(201, model.SuccessResponse(user))
}

// UpdateUser handles updating a user
func UpdateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, model.ErrorResponse(400, "Invalid request"))
		return
	}

	if err := service.UpdateUser(&user); err != nil {
		c.JSON(500, model.ErrorResponse(500, "Failed to update user"))
		return
	}

	c.JSON(200, model.SuccessResponse(user))
}

// DeleteUser handles deleting a user
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := service.DeleteUser(id); err != nil {
		c.JSON(500, model.ErrorResponse(500, "Failed to delete user"))
		return
	}

	c.JSON(204, nil)
}