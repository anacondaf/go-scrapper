package persistence

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type Callback = func(...interface{}) (interface{}, error)

type IRedisService interface {
	Get(ctx context.Context, key string, needMap bool, dto interface{}) (*redis.StringCmd, error)
	Map(mapValue []byte, dto interface{}) error
	GetOrSet(ctx context.Context, key string, expiration time.Duration, dto interface{}, cb Callback) (interface{}, error)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Delete(ctx context.Context, key ...string) error
	HSet(ctx context.Context, key string, setValue ...interface{}) *redis.IntCmd
}
