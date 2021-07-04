package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"mygin/application/logic"
	"mygin/application/models"
	"mygin/settings"
	"mygin/tools/gin_request_response"
	"mygin/tools/ginjwt"
	"strings"
)

//检查accesstoken 是否过期 是否可解析
//单点登录检查 （1）登录后/使用refreshtoken刷新accesstoken之后 都会在redis存入user_id=>accesstoken
//			 （2）在此处每次需要鉴权的时候都检查是否userid和token匹配 如不匹配表面token失效 要求重新登录
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

		//进行redis上绑定的accesstoken对比  单点登录检查
		nowtoken, err := logic.GetOnlineAccesstokenByUserID_ForRelation(mc.User_id)
		if err != nil {
			gin_request_response.Response(c, settings.ErrorAccessTokenSingleLoginCheck, nil)
			c.Abort()
			return
		}
		if nowtoken != parts[1] {
			gin_request_response.Response(c, settings.ErrorAccessTokenSingleLoginHavefreshed, nil)
			c.Abort()
			return
		}

		c.Set(models.ContextUserIdKey, mc.User_id) //为c上下文对象绑定新参数user_id 使用中间件的路由即可获取
		c.Next()

	}
}
