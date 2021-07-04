package middlewares

import (
	"github.com/gin-gonic/gin"
	"mygin/settings"
	"mygin/tools/ginjwt"
	"mygin/tools/ginresponse"
	"strings"
)

//auth middleware
func JwtAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			ginresponse.Response(c, settings.ErrorEmptyToken, nil)
			c.Abort()
			return
		}
		//按照空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			ginresponse.Response(c, settings.ErrorFormatToken, nil)
			c.Abort()
			return
		}

		mc, err := ginjwt.ParseJwtToken(parts[1])
		if err != nil {
			ginresponse.Response(c, settings.ErrorInvalidToken, nil)
			c.Abort()
			return
		}
		c.Set("user_id", mc.User_id) //为c上下文对象绑定新参数user_id 使用中间件的路由即可获取
		c.Next()

	}
}
