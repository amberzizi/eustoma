package agentaction

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

// @title    sendinfo
// @description   访问链接测试
// @auth      amberhu   20210624 15:35
// @param
// @return    err   error   报错
func Sendsqlx(c *gin.Context) {

	//fmt.Println(mysqlset)
	fmt.Println("try connecting")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello sendsendsqlx!",
	})
}
