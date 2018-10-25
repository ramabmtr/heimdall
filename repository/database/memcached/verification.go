package memcached

import (
	"errors"
	"time"

	"github.com/labstack/echo"
	"github.com/ramabmtr/heimdall/config"
	"github.com/ramabmtr/heimdall/helper/formatter"
	"github.com/ramabmtr/heimdall/model"
)

type inMemoryVerificationRepository struct {
	ctx echo.Context
}

func NewVerificationDatabaseRepository(ctx echo.Context) model.VerificationDatabaseRepository {
	return &inMemoryVerificationRepository{
		ctx: ctx,
	}
}

func (c *inMemoryVerificationRepository) Get(key string) (string, error) {
	cacheKey := formatter.CacheKey("OTP", key)
	val, found := config.MemCachedClient.Get(cacheKey)
	if !found {
		return "", nil
	}

	valStr, ok := val.(string)
	if !ok {
		return "", errors.New("cannot cast value to string")
	}

	return valStr, nil
}

func (c *inMemoryVerificationRepository) Store(key, value string) error {
	cacheKey := formatter.CacheKey("OTP", key)
	expTime := time.Duration(config.OtpExpiryTime) * time.Minute

	config.MemCachedClient.Set(cacheKey, value, expTime)

	return nil
}

func (c *inMemoryVerificationRepository) Delete(key string) {
	cacheKey := formatter.CacheKey("OTP", key)
	config.MemCachedClient.Delete(cacheKey)
}
