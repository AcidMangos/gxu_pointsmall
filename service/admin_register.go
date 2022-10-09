package service

import (
	"gxu_pointsmall/dao"
	"gxu_pointsmall/model"
	"gxu_pointsmall/pkg/Err"
	"gxu_pointsmall/pkg/logging"
)

// AdminRegisterService 管理员注册服务
type AdminRegisterService struct {
	UserName        string `form:"user_name" json:"user_name" binding:"required,min=3,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=6,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=6,max=40"`
}

// Valid 验证表单
func (service *AdminRegisterService) Valid() *model.Response {
	var code int
	if service.PasswordConfirm != service.Password {
		code = Err.ERROR_NOT_COMPARE_PASSWORD
		return &model.Response{
			Status: code,
			Msg:    Err.GetMsg(code),
		}
	}

	count := int64(0)
	dao.DBClient.Model(&dao.Admin{}).Where("user_name = ?", service.UserName).Count(&count)
	if count > 0 {
		code = Err.ERROR_EXIST_USER
		return &model.Response{
			Status: code,
			Msg:    Err.GetMsg(code),
		}
	}

	return nil
}

// Register 用户注册
func (service *AdminRegisterService) Register() *model.Response {
	admin := dao.Admin{
		UserName: service.UserName,
	}
	code := Err.SUCCESS
	// 表单验证
	if res := service.Valid(); res != nil {
		return res
	}

	// 加密密码
	if err := admin.SetPassword(service.Password); err != nil {
		logging.Info(err)
		code = Err.ERROR_FAIL_ENCRYPTION
		return &model.Response{
			Status: code,
			Msg:    Err.GetMsg(code),
		}
	}
	// 创建用户
	if err := dao.DBClient.Create(&admin).Error; err != nil {
		logging.Info(err)
		code = Err.ERROR_DATABASE
		return &model.Response{
			Status: code,
			Msg:    Err.GetMsg(code),
		}
	}
	return &model.Response{
		Status: code,
		Msg:    Err.GetMsg(code),
	}
}
