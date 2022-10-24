package service

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gxu_pointsmall/dao"
	"gxu_pointsmall/model"
	"gxu_pointsmall/pkg/Err"
	"gxu_pointsmall/pkg/logging"
	"math/rand"
	"strconv"
	"time"
)

// CreateOrderService 订单创建的服务
type CreateOrderService struct {
	UserID    uint `form:"user_id" json:"user_id"`
	ProductID uint `form:"product_id" json:"product_id"`
	Type      uint `form:"type" json:"type"`
}

func (service *CreateOrderService) Create() model.Response {
	order := dao.Order{
		UserID:    service.UserID,
		ProductID: service.ProductID,
		Type:      service.Type,
	}
	code := Err.SUCCESS
	number := fmt.Sprintf("%09v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000000))
	productNum := strconv.Itoa(int(service.ProductID))
	userNum := strconv.Itoa(int(service.UserID))
	number = number + productNum + userNum
	orderNum, err := strconv.ParseUint(number, 10, 64)
	if err != nil {
		logging.Info(err)
		code = Err.ERROR
		return model.Response{
			Status: code,
			Msg:    Err.GetMsg(code),
			Error:  err.Error(),
		}
	}
	order.OrderNum = orderNum
	//存入数据库
	err = dao.DBClient.Create(&order).Error
	if err != nil {
		logging.Info(err)
		code = Err.ERROR_DATABASE
		return model.Response{
			Status: code,
			Msg:    Err.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//将订单号存入Redis,并设置过期时间
	data := redis.Z{Score: float64(time.Now().Unix()) + 15*time.Minute.Seconds(), Member: orderNum}
	dao.RedisClient.ZAdd(context.Background(), "order", &data)

	return model.Response{
		Status: code,
		Msg:    Err.GetMsg(code),
	}

}
