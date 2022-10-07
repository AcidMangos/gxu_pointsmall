package dao

import "gorm.io/gorm"

type WorkOrder struct {
	gorm.Model
	UserID   uint
	WorkID   uint
	OrderNum uint64
	Price    uint
}
