package config

import (
	"github.com/redis/go-redis/v9"
)

func InitRedis(conf RedisConfig) *redis.Client {
	opts, err := redis.ParseURL(conf.Url)
	if err != nil {
		panic(err)
	}
	rdb := redis.NewClient(opts)

	return rdb
}
