// Package service 包含业务逻辑层（Service 层）
package service

import (
	"fullstack-backend/config"
	"fullstack-backend/model"
	"strings"
)

// MenuNode 树形菜单节点，用于前端侧边栏渲染
// 在 GetMenuTreeByRole 中构建，包含 Children 子节点数组
type MenuNode struct {
	ID       uint        `json:"id"`
	Title    string      `json:"title"`
	Path     string      `json:"path"`
	Icon     string      `json:"icon"`
	ParentID uint        `json:"parent_id"`
	Sort     int         `json:"sort"`
	Hidden   bool        `json:"hidden"`
	Role     string      `json:"role"`
	Children []*MenuNode `json:"children"`
}

// GetMenuList 返回所有菜单（扁平列表）
// 按 Sort 字段升序排列，用于菜单管理页面
func GetMenuList() ([]model.Menu, error) {
	var menus []model.Menu
	if err := config.DB.Order("sort ASC").Find(&menus).Error; err != nil {
		return nil, err
	}
	return menus, nil
}

// GetMenuTreeByRole 根据角色返回树形菜单
// 1. 查询所有未隐藏的菜单
// 2. 根据角色过滤（菜单的 Role 字段包含用户角色才保留）
// 3. 调用 buildMenuTree 递归构建树形结构
func GetMenuTreeByRole(role string) ([]*MenuNode, error) {
	var menus []model.Menu
	if err := config.DB.Where("hidden = ?", false).Order("sort ASC").Find(&menus).Error; err != nil {
		return nil, err
	}

	// 根据角色过滤菜单（菜单 role 字段格式为 "admin,user"）
	var filteredMenus []model.Menu
	for _, menu := range menus {
		if hasRole(menu.Role, role) {
			filteredMenus = append(filteredMenus, menu)
		}
	}

	// 递归构建树形结构，parentID=0 表示根节点
	return buildMenuTree(filteredMenus, 0), nil
}

// hasRole 检查菜单的 role 字段是否包含指定角色
// menuRole: 菜单允许的角色列表，如 "admin,user" 或 "all"
// role: 当前用户角色
// 若菜单未设置角色（空字符串），默认对所有角色可见
func hasRole(menuRole string, role string) bool {
	if menuRole == "" {
		return true
	}
	roles := strings.Split(menuRole, ",")
	for _, r := range roles {
		if strings.TrimSpace(r) == role {
			return true
		}
	}
	return false
}

// buildMenuTree 递归构建菜单树
// menus: 已过滤并排序的菜单列表
// parentID: 当前要构建的父节点 ID，0 表示根节点
// 递归终止条件：没有更多子菜单时返回空切片而非 nil
func buildMenuTree(menus []model.Menu, parentID uint) []*MenuNode {
	var tree []*MenuNode
	for _, menu := range menus {
		if menu.ParentID == parentID {
			node := &MenuNode{
				ID:       menu.ID,
				Title:    menu.Title,
				Path:     menu.Path,
				Icon:     menu.Icon,
				ParentID: menu.ParentID,
				Sort:     menu.Sort,
				Hidden:   menu.Hidden,
				Role:     menu.Role,
				Children: buildMenuTree(menus, menu.ID),
			}
			// 如果没有子菜单，Children 设为空数组而非 nil
			// 确保 JSON 序列化时返回 [] 而非 null
			if node.Children == nil {
				node.Children = []*MenuNode{}
			}
			tree = append(tree, node)
		}
	}
	return tree
}

// CreateMenu 创建新菜单
// 直接插入数据库，由 GORM 自动填充 ID
func CreateMenu(menu *model.Menu) error {
	return config.DB.Create(menu).Error
}

// UpdateMenu 更新菜单
// GORM 的 Save 方法会根据 ID 是否存在自动判断是插入还是更新
func UpdateMenu(menu *model.Menu) error {
	return config.DB.Save(menu).Error
}

// DeleteMenu 软删除菜单
// GORM 软删除：设置 DeletedAt 字段，查询时自动排除已删除记录
func DeleteMenu(id string) error {
	return config.DB.Where("id = ?", id).Delete(&model.Menu{}).Error
}