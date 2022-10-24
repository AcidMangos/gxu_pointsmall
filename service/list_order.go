package service

import (
	"gxu_pointsmall/dao"
	"gxu_pointsmall/model"
	"gxu_pointsmall/pkg/Err"
	"gxu_pointsmall/pkg/logging"
)

// ListOrdersService 订单详情的服务
type ListOrdersService struct {
	Limit int  `form:"limit" json:"limit"`
	Start int  `form:"start" json:"start"`
	Type  uint `form:"type" json:"type" `
}

// List 订单
func (service *ListOrdersService) List(id string) model.Response {
	var orders []dao.Order

	total := int64(0)
	code := Err.SUCCESS
	if service.Limit == 0 {
		service.Limit = 5
	}

	if service.Type == 0 {
		if err := dao.DBClient.Model(&orders).Where("user_id=?", id).Count(&total).Error; err != nil {
			logging.Info(err)
			code = Err.ERROR_DATABASE
			return model.Response{
				Status: code,
				Msg:    Err.GetMsg(code),
				Error:  err.Error(),
			}
		}

		if err := dao.DBClient.Where("user_id=?", id).Limit(service.Limit).Offset(service.Start).Order("created_at desc").Find(&orders).Error; err != nil {
			logging.Info(err)
			code = Err.ERROR_DATABASE
			return model.Response{
				Status: code,
				Msg:    Err.GetMsg(code),
				Error:  err.Error(),
			}
		}
	} else {
		if err := dao.DBClient.Model(&orders).Where("user_id=? AND type=?", id, service.Type).Count(&total).Error; err != nil {
			logging.Info(err)
			code = Err.ERROR_DATABASE
			return model.Response{
				Status: code,
				Msg:    Err.GetMsg(code),
				Error:  err.Error(),
			}
		}

		if err := dao.DBClient.Where("user_id=? AND type=?", id, service.Type).Limit(service.Limit).Offset(service.Start).Order("created_at desc").Find(&orders).Error; err != nil {
			logging.Info(err)
			code = Err.ERROR_DATABASE
			return model.Response{
				Status: code,
				Msg:    Err.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}

	return model.BuildListResponse(model.BuildOrders(orders), uint(total))
}
