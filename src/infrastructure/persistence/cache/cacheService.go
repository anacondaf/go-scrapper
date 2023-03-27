package cache

import (
	"fmt"
	"github.com/kainguyen/go-scrapper/src/config"
	"github.com/redis/go-redis/v9"
)

type CacheService struct {
	config *config.Config
}

func NewCache(config *config.Config) *CacheService {
	return &CacheService{
		config: config,
	}
}

func (c CacheService) CacheConn() (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	fmt.Println("Cache Connect Success!")

	return redisClient, nil
}
