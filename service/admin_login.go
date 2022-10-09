package service

import (
	"github.com/jinzhu/gorm"
	"gxu_pointsmall/dao"
	"gxu_pointsmall/model"
	"gxu_pointsmall/pkg/Err"
	"gxu_pointsmall/pkg/logging"
	"gxu_pointsmall/pkg/util"
)

//AdminLoginService 管理员登录服务
type AdminLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=30"`
	PassWord string `form:"password" json:"password" binding:"required,min=6,max=40"`
}

func (service *AdminLoginService) Login() model.Response {
	var admin dao.Admin
	code := Err.SUCCESS
	if err := dao.DBClient.Where("user_name=?", service.UserName).First(&admin).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			logging.Info(err)
			code = Err.ERROR_NOT_EXIST_USER
			return model.Response{
				Status: code,
				Msg:    Err.GetMsg(code),
			}
		}
		code = Err.ERROR_NOT_EXIST_USER
		return model.Response{
			Status: code,
			Msg:    Err.GetMsg(code),
		}
	}

	if admin.CheckPassword(service.PassWord) == false {
		code = Err.ERROR_NOT_COMPARE
		return model.Response{
			Status: code,
			Msg:    Err.GetMsg(code),
		}

	}
	token, err := util.GenerateToken(service.UserName, service.PassWord, 1)
	if err != nil {
		logging.Info(err)
		code = Err.ERROR_AUTH_TOKEN
		return model.Response{
			Status: code,
			Msg:    Err.GetMsg(code),
		}

	}

	return model.Response{
		Data:   model.TokenData{User: model.BuildAdmin(admin), Token: token},
		Status: code,
		Msg:    Err.GetMsg(code),
	}
}
