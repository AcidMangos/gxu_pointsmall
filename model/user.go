package model

import (
	"gxu_pointsmall/dao"
)

// User 用户序列化器
type User struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}

// BuildUser 序列化用户
func BuildUser(user dao.User) User {
	return User{
		ID:        user.ID,
		UserName:  user.UserName,
		Email:     user.Email,
		Avatar:    user.AvatarURL(),
		CreatedAt: user.CreatedAt.Unix(),
	}
}
