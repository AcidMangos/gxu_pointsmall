package server

import (
	"github.com/gin-gonic/gin"
	"gxu_pointsmall/pkg/Err"
	"gxu_pointsmall/pkg/util"
	"time"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = Err.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = Err.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		if code != Err.SUCCESS {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    Err.GetMsg(code),
				"data":   data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}

// JWTAdmin token验证中间件
func JWTAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = Err.SUCCESS
		token := c.GetHeader("Authorization")
		if token == "" {
			code = Err.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = Err.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = Err.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			} else if claims.Authority == 0 {
				code = Err.ERROR_AUTH_INSUFFICIENT_AUTHORITY
			}
		}

		if code != Err.SUCCESS {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    Err.GetMsg(code),
				"data":   data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
