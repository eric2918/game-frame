package internal

import (
	"frame/conf"

	"github.com/eric2918/leaf/cluster"
)

func handleRpc(id interface{}, f interface{}) {
	cluster.SetRoute(id, ChanRPC)
	skeleton.RegisterChanRPC(id, f)
}

func init() {
	handleRpc("GetGameInfo", GetGameInfo)
}

func GetGameInfo(args []interface{}) ([]interface{}, error) {
	return []interface{}{0, conf.Server.MaxConnNum, conf.Server.TCPAddr, conf.Server.Region}, nil
}
