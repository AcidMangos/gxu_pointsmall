package main

import (
	"gxu_pointsmall/config"
	"gxu_pointsmall/dao"
	"gxu_pointsmall/pkg/logging"
)

func main() {

	// 装载路由
	//r := server.NewRouter()
	//r.Run(":8000")
}

// 数据库初始化连接
func init() {

	// 读取翻译文件
	if err := config.LoadLocales("gxu_pointsmall/config/config.yaml"); err != nil {
		logging.Info(err)
		panic(err)
	}
	mysqlClient, err := dao.NewDBClient()
	redisClient, er := dao.NewRedisClient()
	if err != nil {
		panic(err)
	}
	if er != nil {
		panic(err)
	}
	dao.RedisClient = redisClient
	dao.DBClient = mysqlClient
	dao.Migration()
}
