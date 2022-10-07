package dao

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gxu_pointsmall/config"
)

const mysqlDriver = "mysql"

var DBClient *gorm.DB

func NewDBClient() (*gorm.DB, error) {
	//获取配置单
	mysqlConfig := config.Conf.Storages[mysqlDriver]
	//连接
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/pointsmall?charset=utf8mb4&parseTime=True&loc=Local", mysqlConfig.User, mysqlConfig.Passwd, mysqlConfig.Host)
	//fmt.Printf(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Errorf("mysql connect error:%v", err)
	}
	return db, nil
}
