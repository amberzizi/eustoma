module mygin

go 1.16

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/boombuler/barcode v1.0.1
	github.com/bwmarrin/snowflake v0.3.0 // indirect                    ===雪花ID算法===
	github.com/dchest/captcha v0.0.0-20200903113550-03f5f0333e1f //     ===验证码===
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect         ===jwt===
	github.com/fsnotify/fsnotify v1.4.9 // indirect 					===文件变化监视===
	github.com/gin-gonic/gin v1.7.2 //									===gin框架===
	github.com/go-redis/redis v6.15.9+incompatible // indirect 			===redis连接库===
	github.com/go-sql-driver/mysql v1.6.0 // indirect; daomysql 						===驱动===
	github.com/gohouse/converter v0.0.3 // indirect
	github.com/gohouse/gorose v1.0.5 // indirect  						===dborm取代自有"database/sql"===
	github.com/mattn/go-isatty v0.0.13 // indirect
	github.com/mozillazg/go-pinyin v0.18.0 // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible // indirect  	===日志文件切割===
	github.com/onsi/ginkgo v1.16.4 // indirect
	github.com/onsi/gomega v1.13.0 // indirect
	github.com/spf13/viper v1.8.1 //                                    ===成熟的多类型，支持远程配置中心载入配置库===
	github.com/stretchr/testify v1.7.0
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.7.0 // indirect                           ===文档生成===
	go.uber.org/atomic v1.8.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.17.0 // indirect  								===zap日志库===
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)
