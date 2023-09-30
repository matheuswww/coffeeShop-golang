package redisClient

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

func (m *redisClient) NewRedisConnection() (*redis.Client, error) {
	configErr := m.loadConfigs()
	if configErr != nil {
		return nil,configErr
	}
	ctxTest, cancel := context.WithTimeout(context.Background(), (time.Second * 5))
	defer cancel()
	rdb := redis.NewClient(&redis.Options{
		Addr: m.host,
		Password: m.password,
		DB: m.DB,
	})
	err := rdb.Ping(ctxTest)
	if err.Err() != nil {
		return nil,err.Err()
	}
	return rdb,nil
}

func (m *redisClient) loadConfigs() error {
	mode := os.Getenv("MODE")
	if mode == "PROD" {
		return nil
	} else if mode == "DEV" {
		m.host = "localhost:6379"
		m.password = ""
		m.DB = 0
		return nil
	}
	return errors.New("invalid mod")
}