package main

import (
	"frame/cmd/demo/game"
	"frame/cmd/demo/gate"
	"frame/cmd/demo/login"
	"frame/conf"

	"github.com/eric2918/leaf"
)

func main() {
	// 初始化配置
	conf.Init()

	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)
}
