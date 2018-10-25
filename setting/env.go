package setting

import (
	"strings"

	"github.com/ramabmtr/heimdall/config"
	"github.com/ramabmtr/heimdall/helper/environ"
)

func InitializeEnvVar() {
	config.AppAddress = environ.GetEnv("APP_ADDRESS").Default(":1323").ToString()
	config.AppDebug = environ.GetEnv("APP_DEBUG").Default("1").ToBool()
	config.AppMode = strings.ToLower(environ.GetEnv("APP_MODE").Default("development").ToString())

	config.DatabaseType = strings.ToLower(environ.GetEnv("DATABASE_TYPE").Default("redis").ToString())

	config.OtpExpiryTime = environ.GetEnv("OTP_EXPIRY_TIME").Default("5").ToInt()
}
