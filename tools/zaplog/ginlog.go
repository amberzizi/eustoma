// @Title  ginlog.go
// @Description  zap日志创建，tools.LogerProducter()获取 logger 和 sugarlogger日志对象
// @Author  amberhu  20210625
// @Update
package zaplog

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"mygin/settings"
	"os"
	"time"
)

var logger *zap.Logger
var sugarlogger *zap.SugaredLogger

//type logsetting struct {
//	Log struct{
//		Level  string
//		Maxsize int
//		Maxage int
//		Maxbackups int
//	}
//}

//初始化日志
func InitLogger(logsetting *settings.Log, mode string) {
	var l = new(zapcore.Level)
	err := l.UnmarshalText([]byte(logsetting.Level))
	if err != nil {
		return
	}
	var core zapcore.Core
	if mode == "dev" {
		//开发模式  日志输出到终端
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			//dev开发模式  定义多个日志输出 又记录到文件  又在终端输出
			zapcore.NewCore(getEncoder(), getLogWrite(settings.GetSetting().Log.Maxsize,
				settings.GetSetting().Log.Maxbackups, settings.GetSetting().Log.Maxage, false), l),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(getEncoder(), getLogWrite(settings.GetSetting().Log.Maxsize,
			settings.GetSetting().Log.Maxbackups, settings.GetSetting().Log.Maxage, false), l)
	}
	//zap.AddCaller()增加函数调用信息
	logger = zap.New(core, zap.AddCaller())
	//sugarlogger = logger.Sugar()
	//替换全局logger
	zap.ReplaceGlobals(logger)
	fmt.Printf("Log try init success,use zap.L().debug(...)\n")
}

//logo写入文件  日志每日转储 支持切割
func getLogWrite(maxsize int, maxbackups int, maxage int, ifcompress bool) zapcore.WriteSyncer {
	timeObj := time.Now()
	timestr := timeObj.Format("2006-01-02")
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./logs/" + timestr + ".log",
		MaxSize:    maxsize,    //M
		MaxBackups: maxbackups, //备份数量
		MaxAge:     maxage,     //最多保存多少天
		Compress:   ifcompress, //是否压缩

	}
	//file, _ := os.OpenFile("logs/"+timestr+".log",os.O_APPEND|os.O_CREATE|os.O_RDWR,0744)
	return zapcore.AddSync(lumberJackLogger)
}

//log 格式 json
func getEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

func Printetest() {
	println(settings.GetSetting().Redis.Host)
}
