package server

import (
	"github.com/gin-gonic/gin"
	"gxu_pointsmall/api"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(Cors())
	//store := cookie.NewStore([]byte(sdk.VERSION))
	//r.Use(sessions.Sessions("mysession", store))
	// 路由
	r.GET("/ping", api.Ping)
	//r.POST("/register", api.AdminRegister)
	//r.POST("/login", api.AdminLogin)
	//r.POST("/induct", api.UseRegister)
	r.POST("/creat", api.CreateProduct)
	/*
		v1 := r.Group("/api/v1")
		{
			//注册
			v1.POST("admin/register", api.AdminRegister)

		}*/
	return r
}
