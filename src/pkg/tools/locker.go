package tools

import (
	"frame/pkg/code"
	"frame/pkg/redis"
	"time"

	"github.com/eric2918/leaf/log"
)

const expiration = time.Second * 5

func Locker(key string) int64 {
	ok, err := redis.Client.SetNX(key, 1, expiration).Result()
	if err != nil {
		log.Error("redis链接失败:%s", err.Error())
		return code.ServerInternalError
	}
	if !ok {
		log.Release("操作频繁、请稍候")
		return code.TooManyRequests
	}
	return 200
}

func UnLocker(key string) error {
	err := redis.Client.Del(key).Err()
	if err != nil {
		log.Error("解锁失败:%s", err.Error())
		return err
	}
	return nil
}
