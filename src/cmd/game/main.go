package main

import (
	"frame/cmd/game/center"
	"frame/cmd/game/game"
	"frame/cmd/game/gate"
	"frame/conf"
	"frame/pkg/mongo"
	"frame/pkg/redis"

	"github.com/eric2918/leaf"
)

func main() {
	// 初始化配置
	conf.Init()

	// 初始化mongo
	mongo.Init()

	// 初始化redis
	redis.Init()

	leaf.Run(
		game.Module,
		gate.Module,
		center.Module,
	)
}
