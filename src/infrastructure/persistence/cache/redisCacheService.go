package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kainguyen/go-scrapper/src/config"
	"github.com/kainguyen/go-scrapper/src/core/application/common/persistence"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisCacheService struct {
	redisClient *redis.Client
}

func NewRedisCacheService(config *config.Config) *RedisCacheService {
	redisClient := redis.NewClient(&redis.Options{
		Addr: config.Cache.Address,
		DB:   0,
	})

	fmt.Println("Cache Connect Success!")

	return &RedisCacheService{redisClient: redisClient}
}

func (r RedisCacheService) GetOrSet(ctx context.Context, key string, expiration time.Duration, dto interface{}, cb persistence.Callback) (interface{}, error) {
	mappedValue, err := r.Get(ctx, key, true, dto)

	if err == redis.Nil {
		fmt.Println("Key does not exist. Start write key")

		data, err := cb()
		if err != nil {
			return nil, err
		}

		bytesData, err := json.Marshal(data)

		r.redisClient.Set(ctx, key, bytesData, expiration)

		err = r.Map(bytesData, dto)
		if err != nil {
			return nil, err
		}

		return bytesData, nil
	}

	return mappedValue, err
}

func (r RedisCacheService) Get(ctx context.Context, key string, needMap bool, dto interface{}) (*redis.StringCmd, error) {
	getVal := r.redisClient.Get(ctx, key)

	if needMap && getVal.Err() != redis.Nil {
		byteData, err := getVal.Bytes()
		if err != nil {
			return nil, err
		}

		err = r.Map(byteData, dto)
		if err != nil {
			return nil, err
		}
	}

	return getVal, getVal.Err()
}

func (r RedisCacheService) Map(mapValue []byte, dto interface{}) error {
	if dto == nil {
		return errors.New("cacheQuery: dto interface{} must be provided")
	}

	var err error

	err = json.Unmarshal(mapValue, dto)
	if err != nil {
		return err
	}

	return nil
}

func (r RedisCacheService) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return r.redisClient.Set(ctx, key, value, expiration)
}
