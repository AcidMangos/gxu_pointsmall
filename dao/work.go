package dao

import "gorm.io/gorm"

// Work 商品模型
type Work struct {
	gorm.Model
	WorkID  uint
	Name    string
	Title   string
	Info    string `gorm:"size:1000"`
	ImgPath string
	Price   string
}
