package main

import (
	"frame/cmd/gateway/center"
	"frame/conf"

	"github.com/eric2918/leaf"
)

func main() {
	// 初始化配置
	conf.Init()

	leaf.Run(
		center.Module,
	)
}
