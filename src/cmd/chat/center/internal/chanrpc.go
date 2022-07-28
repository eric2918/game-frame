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
	handleRpc("GetChatInfo", GetChatInfo)
}

func GetChatInfo(args []interface{}) ([]interface{}, error) {
	return []interface{}{0, conf.Server.ListenAddr, conf.Server.Region}, nil
}
