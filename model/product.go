package model

import "gxu_pointsmall/dao"

// Product 商品序列化器
type Product struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Title      string `json:"title"`
	Info       string `json:"info"`
	ImgPath    string `json:"img_path"`
	Price      string `json:"price"`
	CreatedAt  int64  `json:"created_at"`
	ProductNum uint   `json:"num"`
	CategoryId int    `json:"category_id"`
}

// BuildProduct 序列化商品
func BuildProduct(item dao.Product) Product {
	return Product{
		ID:         item.ID,
		Name:       item.Name,
		Title:      item.Title,
		Info:       item.Info,
		ImgPath:    item.ImgPath,
		Price:      item.Price,
		CreatedAt:  item.CreatedAt.Unix(),
		ProductNum: item.ProductNum,
	}
}

// BuildProducts 序列化商品列表
func BuildProducts(items []dao.Product) (products []Product) {
	for _, item := range items {
		product := BuildProduct(item)
		products = append(products, product)
	}
	return products
}
