package main

import (
	"frame/cmd/account/center"
	"frame/cmd/account/gate"
	"frame/cmd/account/login"
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
		center.Module,
		gate.Module,
		login.Module,
	)
}
