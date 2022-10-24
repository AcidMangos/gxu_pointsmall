package service

import (
	"gxu_pointsmall/dao"
	"gxu_pointsmall/model"
	"gxu_pointsmall/pkg/Err"
	"gxu_pointsmall/pkg/logging"
)

// UpdateProductService 商品更新的服务
type UpdateProductService struct {
	ID         uint   `form:"id" json:"id"`
	Name       string `form:"name" json:"name"`
	Title      string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Info       string `form:"info" json:"info" binding:"max=1000"`
	ImgPath    string `form:"img_path" json:"img_path"`
	Price      string `form:"price" json:"price"`
	ProductNum uint   `form:"name" json:"name"`
	categoryId int    `form:"category_id" json:"category_id"`
}

// Update 更新商品
func (service *UpdateProductService) Update() model.Response {
	product := dao.Product{
		Name:       service.Name,
		Title:      service.Title,
		Info:       service.Info,
		ImgPath:    service.ImgPath,
		Price:      service.Price,
		CategoryId: service.categoryId,
	}
	product.ID = service.ID
	code := Err.SUCCESS
	err := dao.DBClient.Save(&product).Error
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
	}
}
