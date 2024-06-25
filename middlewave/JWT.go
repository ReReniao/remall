package middleware

import (
	"github.com/gin-gonic/gin"
	"re-mall/pkg/e"
	"re-mall/pkg/util"
	"re-mall/serializer"
	"time"
)

// JWT 用户登录操作中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := e.Success
		token := c.GetHeader("Authorization")
		if token == "" {
			code = e.InvalidParams
			c.JSON(code, serializer.Response{
				Status: code,
				Data:   e.GetMsg(code),
				Msg:    "缺少token",
				Error:  "",
			})
			c.Abort()
			return
		}
		claims, err := util.ParseToken(token)
		if err != nil {
			code = e.ErrorParseToken

		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ErrorAuthCheckTokenTimeout
		}
		if code != e.Success {
			c.JSON(e.Error, serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  "",
			})
			c.Abort()
			return
		}
		c.Request = c.Request.WithContext(serializer.NewContext(c.Request.Context(), &serializer.UserInfo{
			Id:       claims.ID,
			UserName: claims.UserName,
		}))
		c.Next()
	}
}
