package dao

import "gorm.io/gorm"

// Product 商品模型
type Product struct {
	gorm.Model
	Name    string
	Title   string
	Info    string `gorm:"size:1000"`
	ImgPath string
	Price   string
}
