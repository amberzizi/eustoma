// @Title  agentaction.go
// @Description  测试数据库连接
// @Author  amberhu  20210624
// @Update
package agentaction

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" //init
	"gopkg.in/gcfg.v1"
	"net/http"
)

var db *sql.DB

type mysqlsetting struct {
	//Section struct{
	//	Enabled bool
	//	Path    string
	//}
	Mysql struct {
		Host     string
		Dbname   string
		Username string
		Password string
		Port     string
	}
}

func returnMysqlSetting() *mysqlsetting {
	mysql := mysqlsetting{}
	err := gcfg.ReadFileInto(&mysql, "src/conf/systeminfo.ini")
	if err != nil {
		fmt.Println("Failed to parse config file: %s", err)
	}
	return &mysql
}
func ReturnMsqlDb() *sql.DB {
	mysql := returnMysqlSetting()
	// init daomysql db
	if err := initMySQL(mysql); err != nil {
		fmt.Printf("try connecting fail,err:%v\n", err)
	}
	return db
}

// @title    initMySQL
// @description   初始化数据库连接函数
// @auth      amberhu         20210624 15:35
// @param     daomysql           mysqlsetting     mysql设置参数
// @return    none-db            sql.DB          为全局参数赋值
// @return    err               error           报错
func initMySQL(mysql *mysqlsetting) (err error) {
	dsn := mysql.Mysql.Username + ":" + mysql.Mysql.Password + "@tcp(" + mysql.Mysql.Host + ":" + mysql.Mysql.Port + ")/" + mysql.Mysql.Dbname
	db, err = sql.Open("daomysql", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("try connecting fail,err:%v\n", err)
		return
	}
	//db.SetConnMaxLifetime(time.Second * 10)
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(10)
	return
}

// @title    sendinfo
// @description   访问链接测试
// @auth      amberhu         20210624 15:35
// @param
// @return    err               error           报错
func Sendinfo(c *gin.Context) {
	//	init daomysql setting
	mysql := returnMysqlSetting()

	// init daomysql db
	if err := initMySQL(mysql); err != nil {
		fmt.Printf("try connecting fail,err:%v\n", err)
	}

	//final db close
	defer db.Close()
	fmt.Println("try connecting")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello sendinfoagent!",
	})
}
