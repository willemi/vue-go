package handler

import (
	"fullstack-backend/model"
	"fullstack-backend/service"

	"github.com/gin-gonic/gin"
)

// GetMenuList handles getting menu list
func GetMenuList(c *gin.Context) {
	menus, err := service.GetMenuList()
	if err != nil {
		c.JSON(500, model.ErrorResponse(500, "Failed to get menu list"))
		return
	}

	c.JSON(200, model.SuccessResponse(menus))
}

// CreateMenu handles creating a new menu
func CreateMenu(c *gin.Context) {
	var menu model.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(400, model.ErrorResponse(400, "Invalid request"))
		return
	}

	if err := service.CreateMenu(&menu); err != nil {
		c.JSON(500, model.ErrorResponse(500, "Failed to create menu"))
		return
	}

	c.JSON(201, model.SuccessResponse(menu))
}

// UpdateMenu handles updating a menu
func UpdateMenu(c *gin.Context) {
	var menu model.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(400, model.ErrorResponse(400, "Invalid request"))
		return
	}

	if err := service.UpdateMenu(&menu); err != nil {
		c.JSON(500, model.ErrorResponse(500, "Failed to update menu"))
		return
	}

	c.JSON(200, model.SuccessResponse(menu))
}

// DeleteMenu handles deleting a menu
func DeleteMenu(c *gin.Context) {
	id := c.Param("id")
	if err := service.DeleteMenu(id); err != nil {
		c.JSON(500, model.ErrorResponse(500, "Failed to delete menu"))
		return
	}

	c.JSON(204, nil)
}