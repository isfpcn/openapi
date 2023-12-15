package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

const (
	Offline = iota
	Online
)

type RedisMode int

const (
	RedisModeOfSignal RedisMode = iota
	RedisModeOfSentinel
	RedisModeOfCluster
)

type ServerRegisterInfo struct {
	Name    string
	Ip      string
	Port    int
	Seconds int
	HAPart  string
	HAType  int
}

type Redis struct {
	Client    interface{}
	Addresses []string
	Password  string
	Type      RedisMode
}

var client *Redis

func NewRedisClient(addresses []string, password string, typ int) *Redis {
	client = &Redis{
		Addresses: addresses,
		Password:  password,
		Type:      RedisMode(typ),
	}
	return client
}

func (r *Redis) GetDataByKey(pToken string) ([]byte, error) {
	var data []byte
	var err error
	if r.Type == RedisModeOfCluster {
		data, err = r.Client.(*redis.ClusterClient).Get(pToken).Bytes()
	} else if r.Type == RedisModeOfSignal {
		data, err = r.Client.(*redis.Client).Get(pToken).Bytes()
	}
	return data, err
}

func (r *Redis) KeepAlive(server ServerRegisterInfo) error {
	var err error
	if r.Type == RedisModeOfCluster {
		err = r.NewClientCluster()
	} else if r.Type == RedisModeOfSignal {
		err = r.NewClient()
	}
	if err != nil {
		return err
	}

	//var key string
	//if r.Type == RedisModeOfCluster {
	//	key = fmt.Sprintf("WEB:%s:%s:%d:%s", server.Name, server.Ip, server.Port, server.HAPart)
	//} else if r.Type == RedisModeOfSignal {
	//	key = fmt.Sprintf("WEB:%s:%s:%d", server.Name, server.Ip, server.Port)
	//}
	//
	//now := time.Now()
	//var values = map[string]interface{}{
	//	"service":      server.Name,
	//	"ip":           server.Ip,
	//	"port":         server.Port,
	//	"state":        Online,
	//	"registertime": now.Unix(),
	//	"updatetime":   now.Unix(),
	//}
	//
	//if r.Type == RedisModeOfCluster {
	//	_, err = r.Client.(*redis.ClusterClient).HMSet(key, values).Result()
	//} else if r.Type == RedisModeOfSignal {
	//	_, err = r.Client.(*redis.Client).HMSet(key, values).Result()
	//}
	//if err != nil {
	//	return err
	//}
	//
	//go func() {
	//exit:
	//	for {
	//		select {
	//		case <-time.After(time.Duration(server.Seconds) * time.Second):
	//			// 向redis注册更新状态
	//			now = now.Add(time.Duration(server.Seconds) * time.Second)
	//			if r.Type == RedisModeOfCluster {
	//				_, err = r.Client.(*redis.ClusterClient).HSet(key, "updatetime", now.Unix()).Result()
	//			} else if r.Type == RedisModeOfSignal {
	//				_, err = r.Client.(*redis.Client).HSet(key, "updatetime", now.Unix()).Result()
	//			}
	//			if err != nil {
	//				break exit
	//			}
	//		}
	//	}
	//}()

	return nil
}

// 连接是否断开
func (r *Redis) IsConnection() bool {
	var err error
	if r.Type == RedisModeOfCluster {
		_, err = r.Client.(*redis.ClusterClient).Ping().Result()
	} else if r.Type == RedisModeOfSignal {
		_, err = r.Client.(*redis.Client).Ping().Result()
	}
	if err != nil {
		return false
	}

	return true
}

func (r *Redis) Close() error {
	var err error
	if r.Type == RedisModeOfCluster {
		err = r.Client.(*redis.ClusterClient).Close()
	} else if r.Type == RedisModeOfSignal {
		err = r.Client.(*redis.Client).Close()
	}

	return err
}

func (r *Redis) Get(key string) ([]byte, error) {
	var value []byte
	var err error
	if r.Type == RedisModeOfCluster {
		value, err = r.Client.(*redis.ClusterClient).Get(key).Bytes()
	} else if r.Type == RedisModeOfSignal {
		value, err = r.Client.(*redis.Client).Get(key).Bytes()
	}
	return value, err
}

func (r *Redis) RPush(key string, value string) (int64, error) {
	if r.Type == RedisModeOfCluster {
		return r.Client.(*redis.ClusterClient).RPush(key, value).Result()
	} else if r.Type == RedisModeOfSignal {
		return r.Client.(*redis.Client).RPush(key, value).Result()
	}
	err := fmt.Errorf("Redis RPush Fail")
	return 0, err
}

func (r *Redis) LPop(key string) ([]byte, error) {
	if r.Type == RedisModeOfCluster {
		return r.Client.(*redis.ClusterClient).LPop(key).Bytes()
	} else if r.Type == RedisModeOfSignal {
		return r.Client.(*redis.Client).LPop(key).Bytes()
	}
	err := fmt.Errorf("Redis LPop Fail")
	return []byte(""), err
}

func (r *Redis) Incr(key string) (int64, error) {
	if r.Type == RedisModeOfCluster {
		return r.Client.(*redis.ClusterClient).Incr(key).Result()
	} else if r.Type == RedisModeOfSignal {
		return r.Client.(*redis.Client).Incr(key).Result()
	}
	err := fmt.Errorf("Redis Incr Fail")
	return 0, err
}

func (r *Redis) SetNX(key string, val interface{}, expiration time.Duration) (bool, error) {
	if r.Type == RedisModeOfCluster {
		return r.Client.(*redis.ClusterClient).SetNX(key, val, expiration).Result()
	} else if r.Type == RedisModeOfSignal {
		return r.Client.(*redis.Client).SetNX(key, val, expiration).Result()
	}
	err := fmt.Errorf("Redis SetEx Fail")
	return false, err
}

func (r *Redis) Set(key string, val interface{}, expiration time.Duration) (string, error) {
	if r.Type == RedisModeOfCluster {
		return r.Client.(*redis.ClusterClient).Set(key, val, expiration).Result()
	} else if r.Type == RedisModeOfSignal {
		return r.Client.(*redis.Client).Set(key, val, expiration).Result()
	}
	err := fmt.Errorf("Redis Set Fail")
	return "", err
}

func (r *Redis) Do(args ...interface{}) (interface{}, error) {
	if r.Type == RedisModeOfCluster {
		return r.Client.(*redis.ClusterClient).Do(args...).Result()
	} else if r.Type == RedisModeOfSignal {
		return r.Client.(*redis.Client).Do(args...).Result()
	}
	err := fmt.Errorf("Redis Do Fail")
	return nil, err
}
