package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kainguyen/go-scrapper/src/config"
	"github.com/kainguyen/go-scrapper/src/core/application/common/persistence"
	"github.com/kainguyen/go-scrapper/src/infrastructure/serializer"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"time"
)

type RedisCacheService struct {
	redisClient *redis.Client
	logger      *zerolog.Logger
}

func NewRedisCacheService(config *config.Config, logger *zerolog.Logger) *RedisCacheService {
	redisClient := redis.NewClient(&redis.Options{
		Addr: config.Cache.Address,
		DB:   0,
	})

	logger.Info().Msg("Cache Connect Success")

	return &RedisCacheService{redisClient: redisClient, logger: logger}
}

func (r RedisCacheService) GetOrSet(ctx context.Context, key string, expiration time.Duration, dto interface{}, cb persistence.Callback) (interface{}, error) {
	mappedValue, err := r.Get(ctx, key, true, dto)

	if err == redis.Nil {
		r.logger.Info().Msg("Key does not exist. Start writing key")

		data, err := cb()
		if err != nil {
			return nil, err
		}

		bytesData, err := serializer.Marshal(data)

		if err != nil {
			return nil, err
		}

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

	if !needMap {
		return getVal, getVal.Err()
	}

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
	r.logger.Info().Msg(fmt.Sprintf("CacheService: Set Value For Key %v\n", key))

	bytes, err := serializer.Marshal(value)
	if err != nil {
		return nil
	}

	return r.redisClient.Set(ctx, key, bytes, expiration)
}

func (r RedisCacheService) Delete(ctx context.Context, keys ...string) error {
	pipeline := r.redisClient.Pipeline()

	for _, key := range keys {
		pipeline.Del(ctx, key)
	}

	_, err := pipeline.Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r RedisCacheService) HSet(ctx context.Context, key string, setValue ...interface{}) *redis.IntCmd {
	return r.redisClient.HSet(ctx, key, setValue)
}
