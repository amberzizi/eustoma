module mygin

go 1.16

require (
	github.com/boombuler/barcode v1.0.1
	github.com/fsnotify/fsnotify v1.4.9 // indirect 					===文件变化监视===
	github.com/gin-gonic/gin v1.7.2 //									===gin框架===
	github.com/go-redis/redis v6.15.9+incompatible // indirect 			===redis连接库===
	github.com/go-sql-driver/mysql v1.6.0 //mysql 						===驱动===
	github.com/gohouse/converter v0.0.3 // indirect
	github.com/gohouse/gorose v1.0.5 // indirect  						===dborm取代自有"database/sql"===
	github.com/natefinch/lumberjack v2.0.0+incompatible // indirect  	===日志文件切割===
	github.com/onsi/ginkgo v1.16.4 // indirect
	github.com/onsi/gomega v1.13.0 // indirect
	github.com/spf13/viper v1.8.1 //                                    ===成熟的多类型，支持远程配置中心载入配置库===
	go.uber.org/atomic v1.8.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.17.0 // indirect  								===zap日志库===
	gopkg.in/gcfg.v1 v1.2.3 //ini 										===ini配置文件解析===
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
)
