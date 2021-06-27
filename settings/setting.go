package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//全局配置变量
var SettingGlb *Setting

//同步锁
//var configLock = new(sync.RWMutex)
//

//配置文件结构体
type Setting struct {
	*App   `mapstructure:"App"`
	*Mysql `mapstructure:"Mysql"`
	*Redis `mapstructure:"Redis"`
	*Log   `mapstructure:"Log"`
}
type App struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Runport      string `mapstructure:"runport"`
	Runhost      string `mapstructure:"runhost"`
	Shutdownwait int    `mapstructure:"shutdownwait"`
}
type Mysql struct {
	Host              string `mapstructure:"host"`
	Dbname            string `mapstructure:"dbname"`
	Username          string `mapstructure:"username"`
	Password          string `mapstructure:"password"`
	Port              string `mapstructure:"port"`
	Maxconnection     int    `mapstructure:"maxconnection"`
	Maxidleconnection int    `mapstructure:"maxidleconnection"`
	Prefix            string `mapstructure:"prefix"`
}
type Redis struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
	Poolsize int    `mapstructure:"poolsize"`
}
type Log struct {
	Level      string `mapstructure:"level"`
	Maxsize    int    `mapstructure:"maxsize"`
	Maxage     int    `mapstructure:"maxage"`
	Maxbackups int    `mapstructure:"maxbackups"`
}

//初始化配置文件
func InitSettingViaViper() {
	//载入配置文件
	viper.AddConfigPath("./conf/")
	viper.SetConfigName("systeminfo.ini")
	viper.SetConfigType("ini")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Failed to read config file: %s", err)
	}
	//反序列化到配置结构体
	err = viper.Unmarshal(&SettingGlb)
	if err != nil {
		fmt.Println("Failed to parse config file: %s", err)
	}

	//fsnotify 模块 配置文件变更自动加载
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed: ", e.Name)
		//重新反序列化
		if err := viper.Unmarshal(&SettingGlb); err != nil {
			fmt.Println("Failed to parse config file: %s", err)
		}
	})

}

//获取配置文件参数
func GetSetting() *Setting {
	//configLock.RLock()
	//defer configLock.RUnlock()
	return SettingGlb
}

//载入配置文件 废弃
//使用viper代替
func returnSettingGcfg() {
	//settinginner := Setting{}
	//err := gcfg.ReadFileInto(&settinginner, "conf/systeminfo.ini")
	//if err != nil {
	//	fmt.Println("Failed to parse config file: %s", err)
	//}
	////configLock.Lock()
	//SettingGlb = &settinginner
	////configLock.Unlock()
}

//协程启动定时调用载入配置文件 废弃
//使用fsnotify代替
func freashSetting() {
	//for {
	//	time.Sleep(10 * time.Second)
	//	ReturnSetting()
	//}
}
