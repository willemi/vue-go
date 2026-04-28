package service

import (
	"fullstack-backend/config"
	"fullstack-backend/model"
)

// GetMenuList 返回所有菜单
func GetMenuList() ([]model.Menu, error) {
	var menus []model.Menu
	if err := config.DB.Order("sort ASC").Find(&menus).Error; err != nil {
		return nil, err
	}
	return menus, nil
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