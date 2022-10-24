package dao

import "gorm.io/gorm"

// Order 订单模型
type Order struct {
	gorm.Model
	UserID    uint
	ProductID uint
	OrderNum  uint64
	Price     uint
	Type      uint
}
