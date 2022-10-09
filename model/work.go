package model

// Work 商品模型序列化器
type Work struct {

	Name    string
	Title   string
	Info    string `gorm:"size:1000"`
	ImgPath string
	Price   string
}