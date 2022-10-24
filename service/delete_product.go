package service

import (
	"gxu_pointsmall/dao"
	"gxu_pointsmall/model"
	"gxu_pointsmall/pkg/Err"
	"gxu_pointsmall/pkg/logging"
)

// DeleteProductService 删除商品的服务
type DeleteProductService struct {
}

// Delete 删除商品
func (service *DeleteProductService) Delete(id string) model.Response {
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

	err = dao.DBClient.Delete(&product).Error
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
