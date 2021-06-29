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
	"mygin/tools/qrcode"
	"net/http"
	"strconv"
	"time"
)

/**
* param json
*
*
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
