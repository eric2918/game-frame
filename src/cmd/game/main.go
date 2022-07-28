package main

import (
	"frame/cmd/game/center"
	"frame/cmd/game/game"
	"frame/cmd/game/gate"
	"frame/conf"

	"github.com/eric2918/leaf"
)

func main() {
	// 初始化配置
	conf.Init()

	leaf.Run(
		game.Module,
		gate.Module,
		center.Module,
	)
}
