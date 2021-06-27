package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"math/rand"
	redis3 "mygin/dao/redis"
	"mygin/settings"
	"mygin/tools/qrcode"
	"mygin/tools/snowflake"
	"net/http"
	"strconv"
	"time"
)

func Createuid(c *gin.Context) {
	if err := snowflake.Init(settings.SettingGlb.App.Idstarttime, int64(settings.SettingGlb.App.Machineid)); err != nil {
		fmt.Printf("snowflake init failed,err:%v\n", err)
	} else {
		//zap.L().Info("genid" + strconv.Itoa(int(snowflake.GenId())))
		println(settings.SettingGlb.Idstarttime)
		println(int64(settings.SettingGlb.Machineid))

		println(snowflake.GenId())
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
	rdb := redis3.ReturnRedisDb()
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
