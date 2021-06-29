// @Title  snowflake.go
// @Description  雪花算法id
// @Author  amberhu  20210624
// @Update
package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"mygin/settings"
	"time"
)

var node *snowflake.Node

func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return err
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineID)
	return err
}

func GenId() (int64,error) {
	err := Init(settings.SettingGlb.Idstarttime,int64(settings.SettingGlb.Machineid))
	return node.Generate().Int64(),err
}
