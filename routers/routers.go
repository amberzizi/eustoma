package routers

import (
	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go.uber.org/zap"
	"mygin/application/controllers/community"
	"mygin/application/controllers/post"
	user "mygin/application/controllers/users"
	supiuser "mygin/application/supiapp/controllers/users"
	_ "mygin/docs"
	"mygin/middlewares"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

func helloHandler(c *gin.Context) {
	time.Sleep(time.Second * 5)
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello q1mi!",
	})
}

// SetupRouter 配置路由信息
func SetupRouter(mode string) *gin.Engine {
	//r := gin.Default()
	if mode == gin.ReleaseMode {
		//如果当前模式是发布模式  gin也是
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(GinLogger(zap.L()), GinRecovery(zap.L(), true))

	//doc
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	v1 := r.Group("/api/v1")
	{
		//无须中间件
		//注册
		v1.POST("/signup", user.SignUpHandler) //用户注册
		//登录
		v1.GET("/captcha", user.Captcha)                              //验证码生成
		v1.GET("/captcha/:captchaId", user.GetCaptcha)                //验证码图片获取
		v1.GET("/verify/:captchaId/:value", user.Verify)              //验证码验证
		v1.POST("/login", user.LoginInHandler)                        //用户登录并生成accesstoken+refreshtoken
		v1.POST("/getusernewaccesstoken", user.GetUserNewAccesstoken) //当用户auth认证返回accesstoken401的时候使用此接口用refresstoken获取newaccesstoken
		v1.POST("/getuserinfo", user.GetUserInfer)                    //用户获取用户信息
		//获取社区分类
		v1.GET("/community", community.CommunityHandle)
		v1.GET("/community/:communityId", community.GetCommunityInfoById)
		//帖子
		v1.GET("/postlist/:communityId/:page/:limit", post.PostListHandle)
		//最新全部板块 topn帖子typeid 1 分数最高topn帖子 typeid 2
		v1.GET("/postlistindex/:typeId/:page/:limit", post.PostListIndexHandle)
		//最新社区板块 topn帖子typeid 1 分数最高topn帖子 typeid 2
		v1.GET("/postlistcommunityindex/:communityId/:typeId/:page/:limit", post.PostListCommunityIndexHandle)

		//需要中间件
		//jwt认证中间件
		v1.Use(middlewares.JwtAuthMiddleware())
		{
			//测试登录中间件 用户登录后获取token携带信息
			v1.GET("/getuserinfoafterlogin", user.GetUserInferAfterLogin)
			//帖子详情
			v1.GET("/postlistdetail/:postId", post.PostListDetailHandle)
			//发布帖子
			v1.POST("/postinfo", post.PostInfoHandle)
			//投票
			v1.POST("/vote", post.PostVoteHandle)

		}
	}

	//supi
	supiv1 := r.Group("/api/supiv1")
	{
		//小程序用户注册
		supiv1.POST("/user_center/miniapp_telphone_login", supiuser.User_signup_by_miniapp)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}

// GinLogger 接收gin框架默认的日志
func GinLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		logger.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic
func GinRecovery(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
