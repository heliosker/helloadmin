package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"helloadmin/pkg/app"
	"helloadmin/pkg/errcode"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)
		if s, exist := c.GetQuery("Authorization"); exist {
			token = s
		} else {
			token = c.GetHeader("Authorization")
		}
		if token == "" {
			ecode = errcode.InvalidParams
		} else {
			mc, e := app.ParseToken(token)
			c.Set("username", mc.Username)
			if e != nil {
				switch e.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeOut
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			}
		}
		if ecode != errcode.Success {
			rsp := app.NewResponse(c)
			rsp.Error(ecode)
			c.Abort()
			return
		}
		c.Next()
	}
}
