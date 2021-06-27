package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	mysql2 "mygin/dao/mysql"
	redis2 "mygin/dao/redis"
	routers2 "mygin/routers"
	settings2 "mygin/settings"
	"mygin/tools"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//1.加载配置 使用fsnotify自动加载变化
	//settings.ReturnSetting() ===废弃
	settings2.InitSettingViaViper()
	//2.开启10s刷新配置协程 加载读取锁 ===废弃
	//go settings.FreashSetting()

	//3.加载zaplog
	tools.InitLogger(settings2.SettingGlb.Log)
	//注册 将日志从缓冲区同步给文件
	defer zap.L().Sync()
	zap.L().Debug("logger init success...in main ")

	//4.加载redis初始化检查
	zap.L().Debug(redis2.ReidsInitConnectParamInMain(settings2.SettingGlb.Redis))
	defer redis2.Close()

	//5.加载mysql和mysqlgorose（orm）初始化检查  可关闭其中之一
	//zap.L().Debug(mysql.MysqlInitConnectParamInMain(settings.SettingGlb.Mysql))
	//defer mysql.Close()
	zap.L().Debug(mysql2.MysqlGoroseInitConnectParamInMain(settings2.SettingGlb.Mysql))
	defer mysql2.Gclose()

	//6.载入路由
	r := routers2.SetupRouter()

	//7.协程开机监听端口
	//优雅重启  和 supervisor不可兼得 supervisor会自动拉起监控中的关机进程
	srv := &http.Server{
		Addr:    settings2.SettingGlb.App.Runhost + ":" + settings2.SettingGlb.App.Runport,
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Debug(fmt.Sprint("server listen on..."+settings2.SettingGlb.App.Runport, err))
		}
	}()
	zap.L().Debug(fmt.Sprint("upppppp...", settings2.SettingGlb.App.Runport))

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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(settings2.SettingGlb.App.Shutdownwait))
	defer cancel()
	//5秒内优雅关闭服务器
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Debug(fmt.Sprint("Shutdown server error...", err))
	}
	zap.L().Debug("server exiting...bye")
}
