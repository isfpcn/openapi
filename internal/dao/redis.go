package dao

import (
	"fmt"
	"openapi/internal/config"
	"openapi/internal/dao/redis"
)

func NewRedis() (client *redis.Redis, cf func(), err error) {
	client = redis.NewRedisClient(config.CFG.Redis.Address, config.CFG.Redis.Password, 0)
	err = client.KeepAlive(redis.ServerRegisterInfo{})
	cf = func() {
		fmt.Println("redis close")
		client.Close()
	}
	return
}
