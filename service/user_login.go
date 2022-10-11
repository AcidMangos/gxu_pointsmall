package service

import (
	"github.com/jinzhu/gorm"
	"gxu_pointsmall/dao"
	"gxu_pointsmall/model"
	"gxu_pointsmall/pkg/Err"
	"gxu_pointsmall/pkg/logging"
	"gxu_pointsmall/pkg/util"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=16"`
}

// Login 用户登录函数
func (service *UserLoginService) Login() model.Response {
	var user dao.User
	code := Err.SUCCESS

	if err := dao.DBClient.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		//如果查询不到，返回相应错误
		if gorm.IsRecordNotFoundError(err) {
			logging.Info(err)
			code = Err.ERROR_NOT_EXIST_USER
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

	if user.CheckPassword(service.Password) == false {
		code = Err.ERROR_NOT_COMPARE
		return model.Response{
			Status: code,
			Msg:    Err.GetMsg(code),
		}
	}

	token, err := util.GenerateToken(service.UserName, service.Password, 0)
	if err != nil {
		logging.Info(err)
		code = Err.ERROR_AUTH_TOKEN
		return model.Response{
			Status: code,
			Msg:    Err.GetMsg(code),
		}
	}
	return model.Response{
		Data:   model.TokenData{User: model.BuildUser(user), Token: token},
		Status: code,
		Msg:    Err.GetMsg(code),
	}

}
