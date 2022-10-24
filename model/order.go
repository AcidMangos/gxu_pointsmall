package model

import "gxu_pointsmall/dao"

// Order 订单序列化器
type Order struct {
	ID        uint   `json:"id"`
	OrderNum  uint64 `json:"order_num"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	UserID    uint   `json:"user_id"`
	ProductID uint   `json:"product_id"`
	Name      string `json:"name"`
	ImgPath   string `json:"img_path"`
	Type      uint   `json:"type"`
}

func BuildOrder(item1 dao.Order, item2 dao.Product) Order {
	return Order{
		ID:        item1.ID,
		OrderNum:  item1.OrderNum,
		CreatedAt: item1.CreatedAt.Unix(),
		UpdatedAt: item1.UpdatedAt.Unix(),
		UserID:    item1.UserID,
		ProductID: item1.ProductID,
		Type:      item1.Type,
		Name:      item2.Name,
		ImgPath:   item2.ImgPath,
	}
}

func BuildOrders(items []dao.Order) (orders []Order) {
	for _, item1 := range items {
		item2 := dao.Product{}
		err := dao.DBClient.First(&item2, item1.ProductID).Error
		if err != nil {
			continue
		}
		order := BuildOrder(item1, item2)
		orders = append(orders, order)
	}
	return orders
}
