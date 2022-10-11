package api

import (
	"github.com/gin-gonic/gin"
	"gxu_pointsmall/pkg/logging"
	"gxu_pointsmall/service"
	"path"
)

// UseRegister 用户导入
func UseRegister(c *gin.Context) {
	files, _ := c.FormFile("file")
	dst := path.Join("gxu_pointsmall//upload", files.Filename)
	up := c.SaveUploadedFile(files, dst)
	if up != nil {
		c.JSON(200, ErrorResponse(up))
		logging.Info(up)
	} else {
		res := service.OpenFile(dst)
		c.JSON(200, res)
	}
}
