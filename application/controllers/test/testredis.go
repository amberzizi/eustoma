package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/gohouse/gorose"
	models "mygin/application/models"
	"mygin/dao/daomysql"
	"mygin/dao/daoredis"
	"net/http"
	"time"
)

func Sendredis(c *gin.Context) {

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

func Testq(c *gin.Context) {

	queryGoroseMultiRowDemo(daomysql.ReturnMsqlGoroseConnection())
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello sendinfo!",
	})
}

func queryGoroseMultiRowDemo(connection *gorose.Connection) {
	db := connection.NewSession()
	var user models.User
	var users []models.User
	err2 := db.Table(&user).Select()
	err2 = db.Table(&users).Limit(10).Select()
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println(db.LastSql)
	fmt.Println(user)
	fmt.Println(users)
}
