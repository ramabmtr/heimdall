package redis

import (
	"time"

	depsRedis "github.com/go-redis/redis"
	"github.com/labstack/echo"
	"github.com/ramabmtr/heimdall/config"
	"github.com/ramabmtr/heimdall/helper/formatter"
	"github.com/ramabmtr/heimdall/model"
)

type redisVerificationRepository struct {
	ctx echo.Context
}

func NewVerificationDatabaseRepository(ctx echo.Context) model.VerificationDatabaseRepository {
	return &redisVerificationRepository{
		ctx: ctx,
	}
}

func (c *redisVerificationRepository) Get(key string) (string, error) {
	redisKey := formatter.CacheKey("OTP", key)
	val, err := config.RedisClient.Get(redisKey).Result()
	if err == depsRedis.Nil {
		return "", nil
	}
	if err != nil {
		return "", err
	}

	return val, nil
}

func (c *redisVerificationRepository) Store(key, value string) error {
	redisKey := formatter.CacheKey("OTP", key)
	expTime := time.Duration(config.OtpExpiryTime) * time.Minute

	return config.RedisClient.Set(redisKey, value, expTime).Err()
}

func (c *redisVerificationRepository) Delete(key string) {
	redisKey := formatter.CacheKey("OTP", key)
	config.RedisClient.Del(redisKey)
}
