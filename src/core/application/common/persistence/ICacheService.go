package persistence

import "github.com/redis/go-redis/v9"

type ICacheService interface {
	CacheConn() (*redis.Client, error)
}
