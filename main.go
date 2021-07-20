package main

// @title main
// @version 1.0
// @description eustoma讨论区
// @termsOfService http://swagger.io/terms/

// @contact.name amberhu
// @contact.url https://github.com/amberzizi
// @contact.email amberoracle@163.com

// @license.name Apache 2.0
// @license.url

// @host h
// @BasePath b
import (
	"context"
	"fmt"
	"github.com/shirou/gopsutil/mem"
	"go.uber.org/zap"
	"mygin/dao/daomysql"
	"mygin/dao/daoredis"
	"mygin/routers"
	"mygin/settings"
	"mygin/tools/zaplog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//1.加载配置 使用fsnotify自动加载变化
	//settings.ReturnSetting() ===废弃
	settings.InitSettingViaViper()
	//2.开启10s刷新配置协程 加载读取锁 ===废弃
	//go settings.FreashSetting()

	//3.加载zaplog
	zaplog.InitLogger(settings.SettingGlb.Log, settings.SettingGlb.App.Mode)
	//注册 将日志从缓冲区同步给文件
	defer zap.L().Sync()
	zap.L().Debug("logger init success...in main ")

	//4.加载redis初始化检查
	zap.L().Debug(daoredis.ReidsInitConnectParamInMain(settings.SettingGlb.Redis))
	defer daoredis.Close()

	//5.加载mysql和mysqlgorose（orm）初始化检查  可关闭其中之一
	//zap.L().Debug(daomysql.MysqlInitConnectParamInMain(settings.SettingGlb.Mysql))
	//defer daomysql.Close()
	zap.L().Debug(daomysql.MysqlGoroseInitConnectParamInMain(settings.SettingGlb.Mysql))
	defer daomysql.Gclose()
	//6.载入路由
	r := routers.SetupRouter(settings.SettingGlb.App.Mode)

	//test
	v, _ := mem.VirtualMemory()
	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)
	// convert to JSON. String() is also implemented
	fmt.Println(v)

	//7.协程开机监听端口
	//优雅重启  和 supervisor不可兼得 supervisor会自动拉起监控中的关机进程
	srv := &http.Server{
		Addr:    settings.SettingGlb.App.Runhost + ":" + settings.SettingGlb.App.Runport,
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Debug(fmt.Sprint("server listen on..."+settings.SettingGlb.App.Runport, err))
		}
	}()
	zap.L().Debug(fmt.Sprint("upppppp...linux", settings.SettingGlb.App.Runport))

	//8.平滑优雅关机
	quit := make(chan os.Signal, 1) //创建一个接收信号的通道
	//kill 默认会发送 syscall.SIGTERM 信号  常用的ctrl+c就是触发这种信号
	//kill -2 发送 syscall.SIGINT 信号  添加后才能捕获
	//kill -9 发送 syscall.SIGKILL 信号，不能被捕获 不需添加
	//signal.Notify把接收到的 syscall.SIGINT 或 syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit
	zap.L().Debug("Shutdown server...")
	//创建一个5s超时的context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(settings.SettingGlb.App.Shutdownwait))
	defer cancel()
	//5秒内优雅关闭服务器
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Debug(fmt.Sprint("Shutdown server error...", err))
	}
	zap.L().Debug("server exiting...bye")
}
