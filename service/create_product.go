package service

import (
	"gxu_pointsmall/dao"
	"gxu_pointsmall/model"
	"gxu_pointsmall/pkg/Err"
	"gxu_pointsmall/pkg/logging"
)

// CreateProductService 商品创建的服务
type CreateProductService struct {
	Name       string `form:"name" json:"name"`
	Title      string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Info       string `form:"info" json:"info" binding:"max=1000"`
	ImgPath    string `form:"img_path" json:"img_path"`
	Price      string `form:"price" json:"price"`
	ProductNum uint   `form:"num" json:"num"`
	CategoryId int    `form:"category_id" json:"category_id"`
}

// Create 创建商品
func (service *CreateProductService) Create() model.Response {
	product := dao.Product{
		Name:       service.Name,
		Title:      service.Title,
		Info:       service.Info,
		ImgPath:    service.ImgPath,
		Price:      service.Price,
		ProductNum: service.ProductNum,
		CategoryId: service.CategoryId,
	}
	code := Err.SUCCESS

	err := dao.DBClient.Create(&product).Error
	if err != nil {
		logging.Info(err)
		code = Err.ERROR_DATABASE
		return model.Response{
			Status: code,
			Msg:    Err.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return model.Response{
		Status: code,
		Msg:    Err.GetMsg(code),
		Data:   model.BuildProduct(product),
	}
}
