package middleware

import (
	"github.com/gin-gonic/gin"
	e "helloadmin/pkg/error"
	"helloadmin/pkg/utils"
	"net/http"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			code := e.ERROR_TOKEN_EMPTY
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    code,
				"message": e.Message(code),
			})
			c.Abort()
			return
		}
		mc, ok := utils.ParseToken(authHeader)
		if ok != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    e.ERROR_TOKEN_FAIL,
				"message": e.Message(e.ERROR_TOKEN_FAIL),
			})
			c.Abort()
			return
		}
		c.Set("username", mc.Email)
		c.Next()
	}
}
