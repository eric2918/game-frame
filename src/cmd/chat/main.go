package main

import (
	"frame/cmd/chat/center"
	"frame/cmd/chat/gate"
	"frame/cmd/chat/room"
	"frame/conf"

	"github.com/eric2918/leaf"
)

func main() {
	// 初始化配置
	conf.Init()

	leaf.Run(
		center.Module,
		gate.Module,
		room.Module,
	)
}
