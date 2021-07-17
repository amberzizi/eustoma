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
	"mygin/tools/gin_request_response"
	"mygin/tools/gincaptcha"
	"mygin/tools/qrcode"
	"net/http"
	"strconv"
	"time"
)

//  SignUpHandler  用户注册
//  @Summary  用户注册接口
//  @Description 用户注册接口
//  @Tags 注册
//	@Accept application/json
//	@Produce application/json
//  @Param Authorization header string false "Bearer 用户令牌"
//	@Param object query models.ParamSignUp true "注册参数"
//	@Security ApiKeyAuth
//  @Success 200 {object} []gin_request_response.ResponseData
//  @Router /api/v1/signup [post]
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

//  GetUserInfer  根据userid 获取用户信息
//  @Summary  获取用户信息
//  @Description 获取用户信息
//  @Tags 获取用户信息BY UID
//	@Accept application/json
//	@Produce application/json
//  @Param Authorization header string false "Bearer 用户令牌"
//	@Param object query models.ParamGetuserinfoByUID true "参数"
//	@Security ApiKeyAuth
//  @Success 200 {object} []gin_request_response.ResponseData
//  @Router /api/v1/getuserinf [post]
func GetUserInfer(c *gin.Context) {
	p := new(models.ParamGetuserinfoByUID)
	if err := c.ShouldBindJSON(&p); err != nil {
		//请求参数有误 返回响应  日志记录错误
		zap.L().Error("get userinfo by user_id", zap.Error(err))
		gin_request_response.Response(c, settings.CodeInvalidParam, nil)
		return
	}

	userinfo, err := logic.GetUserInfoByUserId(p)
	if err != nil {
		zap.L().Error("get userinfo by user_id dao failed", zap.Error(err))
	}

	gin_request_response.Response(c, settings.CodeSuccess, userinfo)
}

//生成验证码
func Captcha(c *gin.Context) {
	cap := gincaptcha.GenCaptcha()
	gin_request_response.Response(c, settings.CodeSuccess, cap)
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
		gin_request_response.Response(c, settings.CodeInvalidParam, nil)
		return
	}
	if result {
		gin_request_response.Response(c, settings.CodeSuccess, nil)
		return
	} else {
		gin_request_response.Response(c, settings.CodeVerifyWrong, nil)
		return
	}
}

//  LoginInHandler  用户登录
//  @Summary  用户登录
//  @Description 用户登录
//  @Tags 用户登录
//	@Accept application/json
//	@Produce application/json
//  @Param Authorization header string false "Bearer 用户令牌"
//	@Param object query models.ParamLoginIn true "参数"
//	@Security ApiKeyAuth
//  @Success 200 {object} []gin_request_response.ResponseData
//  @Router /api/v1/login [post]
func LoginInHandler(c *gin.Context) {
	//参数校验
	//var p models.ParamSignUp
	p := new(models.ParamLoginIn)
	if err := c.ShouldBindJSON(&p); err != nil {
		//请求参数有误 返回响应  日志记录错误
		zap.L().Error("LoginIn with invalid param", zap.Error(err))
		gin_request_response.Response(c, settings.CodeInvalidParam, nil)
		return
	}

	//校验参数
	result, err, userinfopublic := logic.LoginCheckPassword(p)
	//用户不存在
	if errors.Is(err, models.ErrorUserNotExist) {
		zap.L().Error("login failed in ErrorUserNotExist", zap.Error(err))
		gin_request_response.Response(c, settings.CodeUserNotExist, nil)
		return
	}
	//用户密码错误
	if errors.Is(err, models.ErrorUserPassowrdWrong) {
		zap.L().Error("login failed in ErrorUserPassowrdWrong", zap.Error(err))
		gin_request_response.Response(c, settings.CodeCheckPasswordWrong, nil)
		return
	}
	//有其他error
	if err != nil {
		zap.L().Error("LoginIn with check password faild", zap.Error(err))
		gin_request_response.Response(c, settings.CodeCheckPasswordThroughWrong, nil)
		return
	}
	//用户名密码匹配 生成jwttoken
	if result {
		jwttoken, err := logic.GenUserJwtToken(p)
		jwttoken_refresh, reerr := logic.GenUserJwtRefreshToken()
		if err != nil || reerr != nil {
			zap.L().Error("Login In  gen jwt token refresh faild", zap.Error(err), zap.Error(reerr))
			gin_request_response.Response(c, settings.ErrorGenToken, nil)
			return
		}
		gin_request_response.Response(c, settings.CodeSuccess,
			map[string]interface{}{"accesstoken": jwttoken, "refreshtoken": jwttoken_refresh, "userinfo": userinfopublic})
		return
	}

	gin_request_response.Response(c, settings.CodePasswordOrUsernameWrong, nil)
	return

}

//测试 登录后获取用户userid
func GetUserInferAfterLogin(c *gin.Context) {
	user_id, err := gin_request_response.GetCurrentUser(c)
	if err != nil {
		gin_request_response.Response(c, settings.ErrorUserNotLogin, nil)
	}
	gin_request_response.Response(c, settings.CodeSuccess, map[string]int64{models.ContextUserIdKey: user_id})
}

func GetUserNewAccesstoken(c *gin.Context) {
	//参数校验
	//var p models.ParamSignUp
	p := new(models.ParamRefreshAccessToken)
	if err := c.ShouldBindJSON(&p); err != nil {
		//请求参数有误 返回响应  日志记录错误
		zap.L().Error("GetUserNewAccesstoken with invalid param", zap.Error(err))
		gin_request_response.Response(c, settings.CodeInvalidParam, nil)
		return
	}

	token, err := logic.GetUserNewAccesstoken(p)
	if err != nil {
		zap.L().Error("GetUserNewAccesstoken  gen jwt token refresh faild", zap.Error(err))
		gin_request_response.Response(c, settings.ErrorGenToken, nil)
		return
	}
	gin_request_response.Response(c, settings.CodeSuccess, map[string]string{"accesstoken": token})
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
