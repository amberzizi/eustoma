// @Title  snowflake.go
// @Description  雪花算法id
// @Author  amberhu  20210624
// @Update
package snowflake

import (
	"github.com/bwmarrin/snowflake"
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

func GetNode() *snowflake.Node {
	return node
}
func GenId() int64 {
	return node.Generate().Int64()
}
