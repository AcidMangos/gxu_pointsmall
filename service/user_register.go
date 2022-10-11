package service

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"gorm.io/gorm"
	"gxu_pointsmall/dao"
	"gxu_pointsmall/model"
	"gxu_pointsmall/pkg/Err"
	"gxu_pointsmall/pkg/logging"
	"strconv"
)

// UserRegisterService 用户注册的服务
type UserRegisterService struct {
	UserName string `binding:"required,min=5,max=15"`
	Password string `binding:"required,min=8,max=16"`
}

var DB *gorm.DB

// Valid 验证表单
func (service *UserRegisterService) Valid() *model.Response {
	var code int

	count := int64(0)
	DB.Model(&dao.User{}).Where("user_name = ?", service.UserName).Count(&count)

	if count > 0 {
		code = Err.ERROR_EXIST_USER
		return &model.Response{
			Status: code,
			Msg:    service.UserName + ":" + Err.GetMsg(code),
		}
	}
	return nil
}

func (service *UserRegisterService) Register(line int) *model.Response {
	user := dao.User{
		UserName: service.UserName,
	}
	code := Err.SUCCESS
	// 验证
	if res := service.Valid(); res != nil {
		return res
	}

	//密码加密
	if err := user.SetPassword(service.Password); err != nil {
		logging.Info(err)
		code = Err.ERROR_FAIL_ENCRYPTION
		return &model.Response{
			Status: code,
			Msg:    strconv.Itoa(line) + ": " + service.UserName + ":" + Err.GetMsg(code),
		}
	}
	// 创建用户
	if err := DB.Create(&user).Error; err != nil {
		logging.Info(err)
		code = Err.ERROR_DATABASE
		return &model.Response{
			Status: code,
			Msg:    strconv.Itoa(line) + ": " + service.UserName + ":" + Err.GetMsg(code),
		}

	}
	return nil
}

func OpenFile(path string) *model.Response {
	code := Err.SUCCESS
	xlsx, err := excelize.OpenFile(path)
	if err != nil {
		code = Err.Error_open_file
		return &model.Response{
			Status: code,
			Msg:    Err.GetMsg(code),
		}
	}
	rows := xlsx.GetRows("Sheet" + "1")

	DB = dao.DBClient.Begin()

	for key, row := range rows {
		if key > 0 {
			user := UserRegisterService{UserName: row[0], Password: row[1]}
			res := user.Register(key)
			if res != nil {
				DB.Rollback()
				return res
			}
		}
	}
	DB.Commit()
	return &model.Response{
		Status: code,
		Msg:    Err.GetMsg(code),
	}
}
