package snowflake

import (
	"frame/conf"
	"log"
	"time"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func init() {
	startTime := conf.Server.StartTime
	machineID := conf.Server.MachineID
	var st time.Time
	st, err := time.Parse("2006-01-02", startTime)
	if err != nil {
		log.Fatal("snowflake parse time error", err)
	}

	snowflake.Epoch = st.UnixNano() / 1e6
	node, err = snowflake.NewNode(machineID)
	if err != nil {
		log.Fatal("snowflake new node error", err)
	}
}

func GenInt() int64 {
	return node.Generate().Int64()
}

func GenStr() string {
	return node.Generate().String()
}
