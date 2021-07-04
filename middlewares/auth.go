package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"mygin/application/models"
	"mygin/settings"
	"mygin/tools/gin_request_response"
	"mygin/tools/ginjwt"
	"strings"
)

//auth middleware
func JwtAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			gin_request_response.Response(c, settings.ErrorEmptyToken, nil)
			c.Abort()
			return
		}
		//按照空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			gin_request_response.Response(c, settings.ErrorFormatToken, nil)
			c.Abort()
			return
		}

		mc, err := ginjwt.ParseJwtToken(parts[1])
		//accesstoken 过期后 返回401
		if errors.Is(err, ginjwt.ErrorAccessTokenExpiredOutOfTime) {
			//zap.L().Error("SignUp failed in logic prase -USER EXSIT", zap.Error(err))
			gin_request_response.Response401(c, settings.ErrorAccessTokenExpiredOutOfTime, nil)
			c.Abort()
			return
		}
		if errors.Is(err, ginjwt.ErrorInvalidToken) {
			//zap.L().Error("SignUp failed in logic prase -USER EXSIT", zap.Error(err))
			gin_request_response.Response401(c, settings.ErrorInvalidToken, nil)
			c.Abort()
			return
		}
		if err != nil {
			gin_request_response.Response(c, settings.ErrorInvalidToken, nil)
			c.Abort()
			return
		}
		c.Set(models.ContextUserIdKey, mc.User_id) //为c上下文对象绑定新参数user_id 使用中间件的路由即可获取
		c.Next()

	}
}
