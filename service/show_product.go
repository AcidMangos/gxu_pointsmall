package service

import (
	"gxu_pointsmall/dao"
	"gxu_pointsmall/model"
	"gxu_pointsmall/pkg/Err"
	"gxu_pointsmall/pkg/logging"
)

// ShowProductService 商品详情的服务
type ShowProductService struct {
}

// Show 商品
func (service *ShowProductService) Show(id string) model.Response {
	var product dao.Product
	code := Err.SUCCESS
	err := dao.DBClient.First(&product, id).Error
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
