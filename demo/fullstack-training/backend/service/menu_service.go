package service

import (
	"fullstack-backend/config"
	"fullstack-backend/model"
	"strings"
)

// MenuNode 树形菜单节点
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

// GetMenuList 返回所有菜单
func GetMenuList() ([]model.Menu, error) {
	var menus []model.Menu
	if err := config.DB.Order("sort ASC").Find(&menus).Error; err != nil {
		return nil, err
	}
	return menus, nil
}

// GetMenuTreeByRole 根据角色返回树形菜单
func GetMenuTreeByRole(role string) ([]*MenuNode, error) {
	var menus []model.Menu
	if err := config.DB.Where("hidden = ?", false).Order("sort ASC").Find(&menus).Error; err != nil {
		return nil, err
	}

	// 根据角色过滤菜单
	var filteredMenus []model.Menu
	for _, menu := range menus {
		if hasRole(menu.Role, role) {
			filteredMenus = append(filteredMenus, menu)
		}
	}

	// 构建树形结构
	return buildMenuTree(filteredMenus, 0), nil
}

// hasRole 检查菜单的 role 字段是否包含指定角色
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
			if node.Children == nil {
				node.Children = []*MenuNode{}
			}
			tree = append(tree, node)
		}
	}
	return tree
}

// CreateMenu 创建新菜单
func CreateMenu(menu *model.Menu) error {
	return config.DB.Create(menu).Error
}

// UpdateMenu 更新菜单
func UpdateMenu(menu *model.Menu) error {
	return config.DB.Save(menu).Error
}

// DeleteMenu 软删除菜单
func DeleteMenu(id string) error {
	return config.DB.Where("id = ?", id).Delete(&model.Menu{}).Error
}
