package redis

import (
	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	config := NewRedisConfig()
	return redis.NewClient(&redis.Options{
		Addr:     config.addr,
		Password: config.passwd,
		DB:       config.db,
	})
}
