package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"math/rand"
	redis3 "mygin/dao/redis"
	"mygin/src/tools"
	"net/http"
	"strconv"
	"time"
)

type configuration struct {
	Enabled  bool
	Path     string
	Username string
	Passwd   string
}

type configuration2 struct {
	Section struct {
		Enabled bool
		Path    string
	}
}

func Sendinfo(c *gin.Context) {
	//json config
	//file, _ := os.Open("src/conf/systeminfo.json")
	//defer file.Close()
	//decoder := json.NewDecoder(file)
	//conf := configuration{}
	//err := decoder.Decode(&conf)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//}
	//fmt.Println(conf)

	//config := configuration2{}
	//err := gcfg.ReadFileInto(&config, "src/conf/systeminfo.ini")
	//if err != nil {
	//	fmt.Println("Failed to parse config file: %s", err)
	//}
	//fmt.Println(config.Section.Path)

	test := false
	if test {
		//测试二维码生成
		randfinal := rand.New(rand.NewSource(time.Now().UnixNano()))
		randname := randfinal.Intn(1000)
		var url = tools.CreateQrcode(200, 200, "testinfo", strconv.Itoa(randname))
		println(url)
	}

	//测试redis
	rdb := redis3.ReturnRedisDb()
	defer rdb.Close()
	//测试watch
	key := "watch_count"
	errw := rdb.Watch(func(tx *redis.Tx) error {
		n, err := tx.Get(key).Int()
		if err != nil && err != redis.Nil {
			fmt.Printf("try connecting fail,err:%v\n", err)
			return err
		}
		println(n)
		time.Sleep(time.Second * 10)
		pipe := tx.Pipeline()
		pipe.Set(key, n+1, 0)
		_, err = pipe.Exec()
		if err != nil {
			fmt.Printf("try connecting fail,err:%v\n", err)
			return err
		}

		println("over")
		return err
	}, key)
	if errw != nil {
		fmt.Printf("try connecting fail,err:%v\n", errw)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello sendinfo!",
	})
}
