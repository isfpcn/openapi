package redis

import (
	"github.com/go-redis/redis"
)

func (r *Redis) NewClient() error {
	opt := &redis.Options{
		Addr:     r.Addresses[0],
		Password: r.Password,
	}
	r.Client = redis.NewClient(opt)

	_, err := r.Client.(*redis.Client).Ping().Result()
	if err != nil {
		return err
	}

	return nil
}
