package redis

import "os"

type redisConfig struct {
	addr   string
	passwd string
	db     int
}

func NewRedisConfig() *redisConfig {

	return &redisConfig{
		addr:   os.Getenv("REDIS_HOST"),
		passwd: os.Getenv("REDIS_PASSWORD"),
		db:     0,
	}
}
