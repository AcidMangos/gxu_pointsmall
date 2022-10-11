package main

import (
	"gxu_pointsmall/config"
	"gxu_pointsmall/dao"
	"gxu_pointsmall/pkg/logging"
	"gxu_pointsmall/server"
)

func main() {

	// 装载路由
	r := server.NewRouter()
	r.Run(":8000")

}

// 数据库初始化连接
func init() {

	// 读取翻译文件
	if err := config.LoadLocales("gxu_pointsmall/config/config.yaml"); err != nil {
		logging.Info(err)
		panic(err)
	}
	mysqlClient, err := dao.NewDBClient()
	if err != nil {
		panic(err)
	}
	dao.DBClient = mysqlClient
	dao.Migration()
}
