package middleware

import (
	"github.com/gin-gonic/gin"
	"helloadmin/pkg/app"
	"helloadmin/pkg/errcode"
	"helloadmin/pkg/utils"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			app.NewResponse(c).Error(errcode.UnauthorizedTokenIsNotExist)
			c.Abort()
			return
		}
		mc, ok := utils.ParseToken(authHeader)
		if ok != nil {
			app.NewResponse(c).Error(errcode.UnauthorizedTokenError)
			c.Abort()
			return
		}
		c.Set("username", mc.Email)
		c.Next()
	}
}
