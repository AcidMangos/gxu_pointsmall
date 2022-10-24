package service

import (
	"gxu_pointsmall/dao"
	"gxu_pointsmall/model"
	"gxu_pointsmall/pkg/Err"
	"gxu_pointsmall/pkg/logging"
)

// SearchProductsService 搜索商品的服务
type SearchProductsService struct {
	Search string `form:"search" json:"search"`
}

// Show 搜索商品
func (service *SearchProductsService) Show() model.Response {
	products := []dao.Product{}
	code := Err.SUCCESS

	err := dao.DBClient.Where("name LIKE ?", "%"+service.Search+"%").Find(&products).Error
	if err != nil {
		logging.Info(err)
		code = Err.ERROR_DATABASE
		return model.Response{
			Status: code,
			Msg:    Err.GetMsg(code),
			Error:  err.Error(),
		}
	}
	products1 := []dao.Product{}
	err = dao.DBClient.Where("info LIKE ?", "%"+service.Search+"%").Find(&products1).Error
	if err != nil {
		logging.Info(err)
		code = Err.ERROR_DATABASE
		return model.Response{
			Status: code,
			Msg:    Err.GetMsg(code),
			Error:  err.Error(),
		}
	}
	products = append(products, products1...)
	return model.Response{
		Status: code,
		Msg:    Err.GetMsg(code),
		Data:   model.BuildProducts(products),
	}
}
