package setting

import (
	"github.com/patrickmn/go-cache"
	"github.com/ramabmtr/heimdall/config"
	"github.com/ramabmtr/heimdall/helper/environ"
	"time"
)

func InitMemCached() error {
	cleanUp := environ.GetEnv("MEMCACHED_CLEANUP_INTERVAL").Default("600000").ToInt()
	cleanUpInterval := time.Duration(cleanUp) * time.Millisecond

	config.MemCachedClient = cache.New(cleanUpInterval, cleanUpInterval)

	return nil
}
