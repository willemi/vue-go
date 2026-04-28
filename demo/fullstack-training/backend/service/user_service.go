package service

import (
	"fullstack-backend/config"
	"fullstack-backend/model"
	"fullstack-backend/utils"

	"gorm.io/gorm"
)

// Login 验证用户并返回 JWT 令牌
func Login(username, password string) (string, model.User, error) {
	var user model.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return "", user, err
	}

	if !utils.CheckPassword(password, user.Password) {
		return "", user, gorm.ErrRecordNotFound
	}

	token, err := utils.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		return "", user, err
	}

	return token, user, nil
}

// GetUserList 返回分页后的用户列表
func GetUserList(username string, page, pageSize int) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	db := config.DB.Model(&model.User{})
	if username != "" {
		db = db.Where("username LIKE ?", "%"+username+"%")
	}

	db.Count(&total)
	if err := db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// CreateUser 创建新用户
func CreateUser(req model.CreateUserRequest) (*model.User, error) {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Username: req.Username,
		Password: hashedPassword,
		Role:     req.Role,
	}

	if err := config.DB.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUser 更新用户
func UpdateUser(req model.UpdateUserRequest) (*model.User, error) {
	var user model.User
	if err := config.DB.First(&user, req.ID).Error; err != nil {
		return nil, err
	}

	// 更新字段
	user.Username = req.Username
	user.Role = req.Role
	if req.Password != "" {
		hashedPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			return nil, err
		}
		user.Password = hashedPassword
	}

	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// DeleteUser 软删除用户
func DeleteUser(id string) error {
	return config.DB.Where("id = ?", id).Delete(&model.User{}).Error
}
