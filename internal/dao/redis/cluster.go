package redis

import (
	"github.com/go-redis/redis"
)

func (r *Redis) NewClientCluster() error {
	opt := &redis.ClusterOptions{
		Addrs:    r.Addresses,
		Password: r.Password,
	}
	r.Client = redis.NewClusterClient(opt)

	_, err := r.Client.(*redis.ClusterClient).Ping().Result()
	if err != nil {
		return err
	}

	return nil
}
