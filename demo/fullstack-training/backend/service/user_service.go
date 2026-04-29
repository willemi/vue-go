// Package service 包含业务逻辑层（Service 层）
// 被 handler 调用，负责处理具体的业务规则，如密码验证、分页查询等
package service

import (
	"fullstack-backend/config"
	"fullstack-backend/model"
	"fullstack-backend/utils"

	"gorm.io/gorm"
)

// Login 验证用户身份并返回 JWT 令牌
// 1. 根据用户名查询数据库
// 2. 使用 bcrypt 验证密码
// 3. 生成 JWT 令牌返回
// 若用户名不存在或密码错误，返回 gorm.ErrRecordNotFound
func Login(username, password string) (string, model.User, error) {
	var user model.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return "", user, err
	}

	// bcrypt.CompareHashAndPassword 会自动处理盐值比较
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
// 支持按用户名模糊搜索（LIKE %username%）
// page 和 pageSize 从请求参数获取，使用 GORM 的 Offset/Limit 实现分页
func GetUserList(username string, page, pageSize int) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	db := config.DB.Model(&model.User{})
	// 按用户名模糊搜索，若 username 为空则跳过此条件
	if username != "" {
		db = db.Where("username LIKE ?", "%"+username+"%")
	}

	// 先计数得到总数，再分页查询
	db.Count(&total)
	if err := db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// CreateUser 创建新用户
// 密码在存入数据库前必须使用 bcrypt 哈希（防止密码明文泄露）
// GORM 会在插入后自动填充 ID 和时间戳字段
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

// UpdateUser 更新用户信息
// 先根据 ID 查找到用户，再更新字段
// password 为可选字段：空字符串表示不修改密码，保留原值
func UpdateUser(req model.UpdateUserRequest) (*model.User, error) {
	var user model.User
	if err := config.DB.First(&user, req.ID).Error; err != nil {
		return nil, err
	}

	// 更新字段
	user.Username = req.Username
	user.Role = req.Role
	// 仅在提供了新密码时才更新（防止空字符串覆盖原密码）
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
// 使用 GORM 的 Delete 方法配合 struct 条件，只设置 DeletedAt 字段为当前时间
func DeleteUser(id string) error {
	return config.DB.Where("id = ?", id).Delete(&model.User{}).Error
}