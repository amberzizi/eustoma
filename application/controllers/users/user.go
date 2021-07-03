package user

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"math/rand"
	"mygin/application/logic"
	"mygin/application/models"
	"mygin/dao/daoredis"
	"mygin/settings"
	"mygin/tools/gincaptcha"
	"mygin/tools/ginresponse"
	"mygin/tools/qrcode"
	"net/http"
	"strconv"
	"time"
)

/**
*用户注册
**/
func SignUpHandler(c *gin.Context) {
	//参数校验
	//var p models.ParamSignUp
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(&p); err != nil {
		//请求参数有误 返回响应  日志记录错误
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		//返回
		ginresponse.Response(c, settings.CodeInvalidParam, nil)
		return
	}

	//注册逻辑
	err := logic.SignUp(p)
	if errors.Is(err, logic.ErrorUserExist) {
		zap.L().Error("SignUp failed in logic prase -USER EXSIT", zap.Error(err))
		ginresponse.Response(c, settings.CodeUserExist, nil)
		return
	}
	if err != nil {
		zap.L().Error("SignUp failed in logic prase", zap.Error(err))
		ginresponse.Response(c, settings.CodeRegisterFail, nil)
		return
	}

	ginresponse.Response(c, settings.CodeSuccess, nil)
}

//根据userid 获取用户信息
func GetUserInfer(c *gin.Context) {
	p := new(models.ParamGetuserinfoByUID)
	if err := c.ShouldBindJSON(&p); err != nil {
		//请求参数有误 返回响应  日志记录错误
		zap.L().Error("get userinfo by user_id", zap.Error(err))
		ginresponse.Response(c, settings.CodeInvalidParam, nil)
		return
	}

	userinfo, err := logic.GetUserInfoByUserId(p)
	if err != nil {
		zap.L().Error("get userinfo by user_id dao failed", zap.Error(err))
	}

	ginresponse.Response(c, settings.CodeSuccess, userinfo)
}

//生成验证码
func Captcha(c *gin.Context) {
	cap := gincaptcha.GenCaptcha()
	ginresponse.Response(c, settings.CodeSuccess, cap)
}

//获取验证码图
func GetCaptcha(c *gin.Context) {
	captchaId := c.Param("captchaId")
	gincaptcha.GetCapid(captchaId, c)
}

//校验验证码
func Verify(c *gin.Context) {
	captchaId := c.Param("captchaId")
	value := c.Param("value")
	result, err := gincaptcha.VerifyCaptcha(captchaId, value)
	if err != nil {
		ginresponse.Response(c, settings.CodeInvalidParam, nil)
		return
	}
	if result {
		ginresponse.Response(c, settings.CodeSuccess, nil)
		return
	} else {
		ginresponse.Response(c, settings.CodeVerifyWrong, nil)
		return
	}
}

/**
*用户登录
完成验证码校验
完成密码校验
*/
func LoginInHandler(c *gin.Context) {
	//参数校验
	//var p models.ParamSignUp
	p := new(models.ParamLoginIn)
	if err := c.ShouldBindJSON(&p); err != nil {
		//请求参数有误 返回响应  日志记录错误
		zap.L().Error("LoginIn with invalid param", zap.Error(err))
		ginresponse.Response(c, settings.CodeInvalidParam, nil)
		return
	}

	//校验参数
	result, err := logic.LoginCheckPassword(p)
	//用户不存在
	if errors.Is(err, logic.ErrorUserNotExist) {
		zap.L().Error("login failed in ErrorUserNotExist", zap.Error(err))
		ginresponse.Response(c, settings.CodeUserNotExist, nil)
		return
	}
	//用户密码错误
	if errors.Is(err, logic.ErrorUserPassowrdWrong) {
		zap.L().Error("login failed in ErrorUserPassowrdWrong", zap.Error(err))
		ginresponse.Response(c, settings.CodeCheckPasswordWrong, nil)
		return
	}
	//有其他error
	if err != nil {
		zap.L().Error("LoginIn with check password faild", zap.Error(err))
		ginresponse.Response(c, settings.CodeCheckPasswordThroughWrong, nil)
		return
	}

	if result {
		ginresponse.Response(c, settings.CodeSuccess, nil)
		return
	}
	ginresponse.Response(c, settings.CodePasswordOrUsernameWrong, nil)
	return

}

func Sendinfo(c *gin.Context) {

	test := false
	if test {
		//测试二维码生成
		randfinal := rand.New(rand.NewSource(time.Now().UnixNano()))
		randname := randfinal.Intn(1000)
		var url = qrcode.CreateQrcode(200, 200, "testinfo", strconv.Itoa(randname))
		println(url)
	}

	//测试redis
	rdb := daoredis.ReturnRedisDb()
	defer rdb.Close()
	//测试watch
	key := "watch_count"
	errw := rdb.Watch(func(tx *redis.Tx) error {
		n, err := tx.Get(key).Int()
		if err != nil && err != redis.Nil {
			fmt.Printf("try connecting fail,err:%v\n", err)
			return err
		}
		println(n)
		time.Sleep(time.Second * 10)
		pipe := tx.Pipeline()
		pipe.Set(key, n+1, 0)
		_, err = pipe.Exec()
		if err != nil {
			fmt.Printf("try connecting fail,err:%v\n", err)
			return err
		}

		println("over")
		return err
	}, key)
	if errw != nil {
		fmt.Printf("try connecting fail,err:%v\n", errw)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello sendinfo!",
	})
}
