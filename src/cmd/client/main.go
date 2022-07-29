package main

import (
	"frame/cmd/client/client"
	"frame/conf"
	"frame/msg"

	"github.com/eric2918/leaf"
)

func main() {
	// 初始化配置
	conf.Init()

	client.Init(msg.Processor)

	leaf.Run(
		client.Module,
	)
}
