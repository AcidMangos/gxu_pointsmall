package api

import (
	"github.com/gin-gonic/gin"
	"gxu_pointsmall/pkg/logging"
	"gxu_pointsmall/service"
)

// CreateProduct 创建商品
func CreateProduct(c *gin.Context) {
	service := service.CreateProductService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

// UpdateProduct 更新商品信息
func UpdateProduct(c *gin.Context) {
	service := service.UpdateProductService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}

}

// DeleteProduct 删除商品的接口
func DeleteProduct(c *gin.Context) {
	service := service.DeleteProductService{}
	res := service.Delete(c.Param("id"))
	c.JSON(200, res)
}

// SearchProducts 搜索商品的接口
func SearchProducts(c *gin.Context) {
	service := service.SearchProductsService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Show()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
