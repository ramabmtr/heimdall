package main

import (
	"github.com/ramabmtr/heimdall/api"
	"github.com/ramabmtr/heimdall/config"
	"github.com/ramabmtr/heimdall/helper/log"
	"github.com/ramabmtr/heimdall/setting"
)

func main() {
	setting.InitializeEnvVar()
	setting.InitializeLogger()

	if config.VerificationServiceType == "" {
		log.GetLogger(nil).Fatal("SERVICE_TYPE not set!")
	}

	switch config.VerificationDatabaseType {
	case config.DBTypeMemCached:
		if err := setting.InitMemCached(); err != nil {
			log.GetLogger(nil).WithError(err).Fatal("fail to connect to memcached")
		}
	case config.DBTypeRedis:
		fallthrough
	default:
		if err := setting.InitRedis(); err != nil {
			log.GetLogger(nil).WithError(err).Fatal("fail to connect to redis")
		}
	}

	api.Run()
}
