package service

import (
	"fullstack-backend/config"
	"fullstack-backend/model"
	"fullstack-backend/utils"

	"gorm.io/gorm"
)

// Login authenticates a user and returns a JWT token
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

// GetUserList returns a paginated list of users
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

// CreateUser creates a new user
func CreateUser(user *model.User) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return config.DB.Create(user).Error
}

// UpdateUser updates a user
func UpdateUser(user *model.User) error {
	if user.Password != "" {
		hashedPassword, err := utils.HashPassword(user.Password)
		if err != nil {
			return err
		}
		user.Password = hashedPassword
	}
	return config.DB.Save(user).Error
}

// DeleteUser soft-deletes a user
func DeleteUser(id string) error {
	return config.DB.Where("id = ?", id).Delete(&model.User{}).Error
}