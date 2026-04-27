package service

import (
	"fullstack-backend/config"
	"fullstack-backend/model"
)

// GetMenuList returns all menus
func GetMenuList() ([]model.Menu, error) {
	var menus []model.Menu
	if err := config.DB.Order("sort ASC").Find(&menus).Error; err != nil {
		return nil, err
	}
	return menus, nil
}

// CreateMenu creates a new menu
func CreateMenu(menu *model.Menu) error {
	return config.DB.Create(menu).Error
}

// UpdateMenu updates a menu
func UpdateMenu(menu *model.Menu) error {
	return config.DB.Save(menu).Error
}

// DeleteMenu soft-deletes a menu
func DeleteMenu(id string) error {
	return config.DB.Where("id = ?", id).Delete(&model.Menu{}).Error
}