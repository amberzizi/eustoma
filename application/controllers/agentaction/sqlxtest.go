package agentaction

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"mygin/application/models"
	"net/http"
)

// @title    sendinfo
// @description   访问链接测试
// @auth      amberhu   20210624 15:35
// @param
// @return    err   error   报错
func Sendsqlx(c *gin.Context) {
	//获取数据库初始化的连接对象
	db_sqx := ReturnMsqlDb()
	queryRowDemo2(db_sqx)
	//fmt.Println(mysqlset)
	fmt.Println("try connecting")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello sendsendsqlx!",
	})
}

func queryRowDemo2(db *sql.DB) {

	sqlStr := "select id,name,age from user where id=?"
	var u myginuser.User
	err := db.QueryRow(sqlStr, 1).Scan(&u.Id, &u.Name, &u.Age)
	if err != nil {
		fmt.Printf("scan failed err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.Id, u.Name, u.Age)
}
