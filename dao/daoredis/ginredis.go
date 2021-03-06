// @Title  ginredis.go
// @Description  zap日志创建，tools.ReturnRedisDb() redis对象  单点 未扩展哨兵 集群
// @Author  amberhu  20210625
// @Update

//测试redis
//rdb := tools.ReturnRedisDb()
//defer rdb.Close()   回收*****
package daoredis

import (
	"fmt"
	"github.com/go-redis/redis"
	"mygin/settings"
)

var rdb *redis.Client

//var Rdb *daoredis.Client

//对外返回redis连接对象
//可以直接用redis.Rdb
func ReturnRedisDb() *redis.Client {
	return rdb
}

//初始化redis 连接
func initRedisClient(redisset *settings.Redis) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     redisset.Host,
		Password: redisset.Password,
		DB:       redisset.Db,
		PoolSize: redisset.Poolsize,
	})

	_, err = rdb.Ping().Result()
	return err
}

//main里面用的初始化参数文件 初始化连接
func ReidsInitConnectParamInMain(redisset *settings.Redis) string {
	err := initRedisClient(redisset)
	if err != nil {
		fmt.Printf("daoredis try connecting fail,err:%v\n", err)
		return "daoredis try connecting fail,err"
	} else {
		return "daoredis try connecting success"
	}
}

func Close() {
	_ = rdb.Close()
}
