package config

import (
	"github.com/go-redis/redis"
	"github.com/patrickmn/go-cache"
)

var (
	AppAddress string
	AppDebug   bool
	AppMode    string

	DatabaseType string

	OtpExpiryTime int

	RedisClient     *redis.Client
	MemCachedClient *cache.Cache
)
