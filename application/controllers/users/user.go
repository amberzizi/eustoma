package user

import (
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
		c.JSON(http.StatusOK, gin.H{
			"code": "1002",
			"msg":  settings.CodeSetting[1002],
		})
		return
	}

	//注册逻辑
	err := logic.SignUp(p)
	if err != nil {
		zap.L().Error("SignUp failed in logic prase", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": "1004",
			"msg":  settings.CodeSetting[1004],
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": "1001",
		"msg":  settings.CodeSetting[1001],
	})
}

//根据userid 获取用户信息
func GetUserInfer(c *gin.Context) {
	p := new(models.ParamGetuserinfoByUID)
	if err := c.ShouldBindJSON(&p); err != nil {
		//请求参数有误 返回响应  日志记录错误
		zap.L().Error("get userinfo by user_id", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": "1002",
			"msg":  settings.CodeSetting[1002],
		})
		return
	}

	userinfo, err := logic.GetUserInfoByUserId(p)
	if err != nil {
		zap.L().Error("get userinfo by user_id dao failed", zap.Error(err))
	}

	c.JSON(http.StatusOK, gin.H{
		"code": "1001",
		"msg":  settings.CodeSetting[1001],
		"data": userinfo,
	})
}

//生成验证码
func Captcha(c *gin.Context) {
	cap := gincaptcha.GenCaptcha()
	c.JSON(http.StatusOK, gin.H{
		"code": "1001",
		"msg":  settings.CodeSetting[1001],
		"data": cap,
	})
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
		c.JSON(http.StatusOK, gin.H{
			"code": "1002",
			"msg":  settings.CodeSetting[1002],
			"data": [1]string{},
		})
		return
	}
	if result {
		c.JSON(http.StatusOK, gin.H{
			"code": "1001",
			"msg":  settings.CodeSetting[1001],
			"data": [1]string{},
		})
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
		c.JSON(http.StatusOK, gin.H{
			"code": "1002",
			"msg":  settings.CodeSetting[1002],
		})
		return
	}

	//校验参数
	result, err := logic.LoginCheckPassword(p)
	if err != nil {
		zap.L().Error("LoginIn with check password faild", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"code": "1005",
			"msg":  settings.CodeSetting[1005],
		})
		return
	}

	if !result {
		c.JSON(http.StatusOK, gin.H{
			"code": "1006",
			"msg":  settings.CodeSetting[1006],
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": "1001",
			"msg":  settings.CodeSetting[1001],
		})
	}

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
