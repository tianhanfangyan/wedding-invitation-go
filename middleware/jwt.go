package middleware

import (
	"github.com/gin-gonic/gin"
	"wedding-invitation-go/models"
	"wedding-invitation-go/utils"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			models.RetFailWithDetail(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}

		_, err := utils.ParseToken(token)
		if err != nil {
			models.RetFailWithDetail(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}

		// 处理请求
		c.Next()
	}
}
