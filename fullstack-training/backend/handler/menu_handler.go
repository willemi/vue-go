// Package handler 包含所有 HTTP 请求处理函数（Controller 层）
package handler

import (
	"fullstack-backend/model"
	"fullstack-backend/service"

	"github.com/gin-gonic/gin"
)

// GetMenuList 处理获取菜单列表（扁平结构，用于菜单管理页面）
// GET /api/menu/list
// 返回所有菜单记录（含已隐藏的），无层级关系，用于管理员编辑
// 响应：{ code: 200, data: menus[] }
func GetMenuList(c *gin.Context) {
	menus, err := service.GetMenuList()
	if err != nil {
		c.JSON(500, model.ErrorResponse(500, "Failed to get menu list"))
		return
	}

	c.JSON(200, model.SuccessResponse(menus))
}

// GetMenuTree 处理获取当前用户的菜单树（用于侧边栏导航）
// GET /api/menu/tree
// 根据登录用户的角色过滤菜单，并构建树形层级结构返回
// role 从 AuthMiddleware 注入的 context 中获取
// 响应：{ code: 200, data: treeNodes[] }
func GetMenuTree(c *gin.Context) {
	// 从 context 中获取用户角色（由 AuthMiddleware 设置）
	role, _ := c.Get("role")

	menus, err := service.GetMenuTreeByRole(role.(string))
	if err != nil {
		c.JSON(500, model.ErrorResponse(500, "Failed to get menu tree"))
		return
	}

	c.JSON(200, model.SuccessResponse(menus))
}

// CreateMenu 处理创建新菜单
// POST /api/menu/add
// 请求体：{ title, path, icon?, parent_id?, sort?, hidden?, role? }
// 响应：{ code: 201, data: menu }
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

// UpdateMenu 处理更新菜单
// PUT /api/menu/edit
// 请求体：{ id, title, path, icon?, parent_id?, sort?, hidden?, role? }
// 响应：{ code: 200, data: menu }
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

// DeleteMenu 处理删除菜单（软删除）
// DELETE /api/menu/delete/:id
// 响应：{ code: 204, data: null }
func DeleteMenu(c *gin.Context) {
	id := c.Param("id")
	if err := service.DeleteMenu(id); err != nil {
		c.JSON(500, model.ErrorResponse(500, "Failed to delete menu"))
		return
	}

	c.JSON(204, nil)
}