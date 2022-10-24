package service

import (
	"github.com/jinzhu/gorm"
	"gxu_pointsmall/dao"
	"gxu_pointsmall/model"
	"gxu_pointsmall/pkg/Err"
	"gxu_pointsmall/pkg/logging"
)

// ShowOrderService 订单详情的服务
type ShowOrderService struct {
}

// Show 订单
func (service *ShowOrderService) Show(num string) model.Response {
	var order dao.Order
	var product dao.Product
	code := Err.SUCCESS
	//根据id查找order
	if err := dao.DBClient.Where("order_num=?", num).First(&order).Error; err != nil {
		logging.Info(err)
		code = Err.ERROR_DATABASE
		return model.Response{
			Status: code,
			Msg:    Err.GetMsg(code),
		}
	}
	//根据order查找product
	if err := dao.DBClient.Where("id=?", order.ProductID).First(&product).Error; err != nil {
		//如果查询不到，返回相应错误
		if gorm.IsRecordNotFoundError(err) {
			logging.Info(err)
			code = Err.ERROR_NOT_EXIST_PRODUCT
			return model.Response{
				Status: code,
				Msg:    Err.GetMsg(code),
			}
		}
		logging.Info(err)
		code = Err.ERROR_DATABASE
		return model.Response{
			Status: code,
			Msg:    Err.GetMsg(code),
		}
	}

	return model.Response{
		Status: code,
		Msg:    Err.GetMsg(code),
		Data:   model.BuildOrder(order, product),
	}
}
