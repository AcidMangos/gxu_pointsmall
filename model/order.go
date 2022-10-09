package model

// Order 订单序列化器
type Order struct {
	ID            uint   `json:"id"`
	OrderNum      uint64 `json:"order_num"`
	CreatedAt     int64  `json:"created_at"`
	UpdatedAt     int64  `json:"updated_at"`
	UserID        uint   `json:"user_id"`
	ProductID     uint   `json:"product_id"`
	Num           uint   `json:"num"`
	Name          string `json:"name"`
	ImgPath       string `json:"img_path"`
}