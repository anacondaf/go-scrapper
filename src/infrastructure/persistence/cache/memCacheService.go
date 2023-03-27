package cache

import "github.com/kainguyen/go-scrapper/src/config"

type MemCacheService struct {
}

func NewMemCacheService(config *config.Config) *MemCacheService {
	return &MemCacheService{}
}
