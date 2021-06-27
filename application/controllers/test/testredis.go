package test

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/gohouse/gorose"
	myginuser2 "mygin/application/models"
	mysql2 "mygin/dao/mysql"
	redis3 "mygin/dao/redis"
	"net/http"
	"time"
)

func Sendredis(c *gin.Context) {

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

func Testq(c *gin.Context) {
	queryMultiRowDemo(mysql2.ReturnMsqlDb())
	queryGoroseMultiRowDemo(mysql2.ReturnMsqlGoroseConnection())
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello sendinfo!",
	})
}
func queryMultiRowDemo(db *sql.DB) {
	sqlStr := "select id, name, age from user where id > ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var u myginuser2.User
		err := rows.Scan(&u.Id, &u.Name, &u.Age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.Id, u.Name, u.Age)
	}
}

func queryGoroseMultiRowDemo(connection *gorose.Connection) {
	db := connection.NewSession()
	var user myginuser2.User
	var users []myginuser2.User
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
