// Package utils 提供加密、签名等工具函数
// 当前包括 JWT 令牌生成与解析、密码哈希与验证
package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims JWT 令牌中存储的用户身份信息
// 包含用户 ID、用户名、角色，以及 JWT 标准声明（过期时间等）
type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// jwtSecret 用于签名和验证 JWT 令牌的密钥
// 生产环境中应从环境变量或安全存储中读取
var jwtSecret = []byte("fullstack-secret-key-2024")

// GenerateToken 生成 JWT 令牌
// userID: 用户唯一标识
// username: 用户名
// role: 用户角色（admin/user）
// 返回: 签名后的 token 字符串，24 小时后过期
func GenerateToken(userID uint, username, role string) (string, error) {
	claims := Claims{
		UserID:   userID,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ParseToken 解析并验证 JWT 令牌
// tokenString: JWT 令牌字符串
// 返回: 解析后的 Claims（包含用户信息），若令牌无效或已过期则返回错误
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}