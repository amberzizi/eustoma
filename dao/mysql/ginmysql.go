package mysql

import (
	"database/sql"
	"fmt"
	"github.com/gohouse/gorose"
	"go.uber.org/zap"
	settings2 "mygin/settings"
)

//var Db *sql.DB
var db *sql.DB

//var Gdb *gorose.Connection
var gdb *gorose.Connection

//获取DB对象
//直接可以用mysq.Db获取
func ReturnMsqlDb() *sql.DB {
	return db
}
func ReturnMsqlGoroseConnection() *gorose.Connection {
	return gdb
}

// @title    initMySQL
// @description   初始化数据库连接函数
// @auth      amberhu         20210624 15:35
// @param     mysql           mysqlsetting     mysql设置参数
// @return    none-db            sql.DB          为全局参数赋值
// @return    err               error           报错
func initMySQL(mysqlset *settings2.Mysql) (err error) {
	dsn := mysqlset.Username + ":" + mysqlset.Password + "@tcp(" + mysqlset.Host + ":" + mysqlset.Port + ")/" + mysqlset.Dbname
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		zap.L().Error("mysql init faild", zap.Error(err))
	}

	err = db.Ping()
	if err != nil {
		zap.L().Error("mysql init ping faild", zap.Error(err))
	}
	//db.SetConnMaxLifetime(time.Second * 10)
	db.SetMaxOpenConns(mysqlset.Maxconnection)
	db.SetMaxIdleConns(mysqlset.Maxidleconnection)
	return err
}

func initGoroseMySQL(mysqlset *settings2.Mysql) (err error) {
	var dbConfig1 = &gorose.DbConfigSingle{
		Driver:          "mysql",                                                                                                              // 驱动: mysql/sqlite/oracle/mssql/postgres
		EnableQueryLog:  true,                                                                                                                 // 是否开启sql日志
		SetMaxOpenConns: mysqlset.Maxconnection,                                                                                               // (连接池)最大打开的连接数，默认值为0表示不限制
		SetMaxIdleConns: mysqlset.Maxidleconnection,                                                                                           // (连接池)闲置的连接数
		Prefix:          "",                                                                                                                   // 表前缀
		Dsn:             mysqlset.Username + ":" + mysqlset.Password + "@tcp(" + mysqlset.Host + ":" + mysqlset.Port + ")/" + mysqlset.Dbname, // 数据库链接
	}
	gdb, err = gorose.Open(dbConfig1)
	if err != nil {
		zap.L().Error("mysql gorose init faild", zap.Error(err))
		return
	}
	return err
}

//main里面用的初始化参数文件
func MysqlInitConnectParamInMain(mysqlset *settings2.Mysql) string {
	err := initMySQL(mysqlset)
	if err != nil {
		fmt.Printf("mysql try connecting fail,err:%v\n", err)
		return "mysql try connecting fail"
	} else {
		fmt.Printf("mysql try connecting success\n")
		return "mysql try connecting success"
	}
}

//main里面用的初始化参数文件
func MysqlGoroseInitConnectParamInMain(mysqlset *settings2.Mysql) string {
	err := initGoroseMySQL(mysqlset)
	if err != nil {
		fmt.Printf("mysql Gorose try connecting fail,err:%v\n", err)
		return "mysql Gorose try connecting fail"
	} else {
		fmt.Printf("mysql Gorose try connecting success\n")
		return "mysql Gorose try connecting success"
	}
}

func Close() {
	_ = db.Close()
}

func Gclose() {
	_ = gdb.Close()
}
