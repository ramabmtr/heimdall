package setting

import (
	"github.com/go-redis/redis"
	"github.com/ramabmtr/heimdall/config"
	"github.com/ramabmtr/heimdall/helper/environ"
)

func InitRedis() error {
	url := environ.GetEnv("REDIS_URL").Default("localhost:6379").ToString()
	pass := environ.GetEnv("REDIS_PASSWORD").Default("").ToString()
	db := environ.GetEnv("REDIS_DB").Default("0").ToInt()

	config.RedisClient = redis.NewClient(&redis.Options{
		Addr:     url,
		Password: pass,
		DB:       db,
	})

	if _, err := config.RedisClient.Ping().Result(); err != nil {
		return err
	}

	return nil
}
