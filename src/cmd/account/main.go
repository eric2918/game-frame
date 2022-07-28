package main

import (
	"frame/cmd/account/center"
	"frame/cmd/account/gate"
	"frame/cmd/account/login"
	"frame/conf"

	"github.com/eric2918/leaf"
)

func main() {
	// 初始化配置
	conf.Init()

	leaf.Run(
		center.Module,
		gate.Module,
		login.Module,
	)
}
