package main

import "gxu_pointsmall/dao"

func main() {

}

// 数据库初始化连接
func init() {
	mysqlClient, err := dao.NewDBClient()
	if err != nil {
		panic(err)
	}
	dao.DBClient = mysqlClient
}
