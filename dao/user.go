package dao

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UseId          uint `gorm:"unique"`
	UserName       string
	Email          string `gorm:"unique"`
	PasswordDigest string
	Points         uint `gorm:"default:0"`
}

// GetUser 获取用户
func GetUser(id interface{}) (User, error) {
	var user User
	res := DBClient.First(&user, id)
	return user, res.Error
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	pw, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(pw)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err != nil
}
