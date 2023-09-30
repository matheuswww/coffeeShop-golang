package redisClient

import "github.com/redis/go-redis/v9"

func NewRedis() Redis {
	return &redisClient{}
}

type Redis interface {
	NewRedisConnection() (*redis.Client, error)
}

type redisClient struct {
	host     string
	password string
	DB     int
}
