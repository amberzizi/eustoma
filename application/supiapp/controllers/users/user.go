package users

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mygin/application/logic"
	"mygin/application/models"
	"mygin/settings"
	"mygin/tools/gin_request_response"
	"net/http"
)

func SignUpHandler(c *gin.Context) {
	//参数校验
	//var p models.ParamSignUp
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(&p); err != nil {
		//请求参数有误 返回响应  日志记录错误
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		//返回
		gin_request_response.Response(c, settings.CodeInvalidParam, nil)
		return
	}

	//注册逻辑
	err := logic.SignUp(p)
	if errors.Is(err, models.ErrorUserExist) {
		zap.L().Error("SignUp failed in logic prase -USER EXSIT", zap.Error(err))
		gin_request_response.Response(c, settings.CodeUserExist, nil)
		return
	}
	if err != nil {
		zap.L().Error("SignUp failed in logic prase", zap.Error(err))
		gin_request_response.Response(c, settings.CodeRegisterFail, nil)
		return
	}

	gin_request_response.Response(c, settings.CodeSuccess, nil)
}

//小程序用户注册
func User_signup_by_miniapp(c *gin.Context) {

	p := new(models.ParamUserMiniappSignUp)
	if err := c.ShouldBindJSON(&p); err != nil {
		//请求参数有误 返回响应  日志记录错误
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		//返回
		gin_request_response.Response(c, settings.CodeInvalidParam, nil)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello q1mi322323!",
	})
}
