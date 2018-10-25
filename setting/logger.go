package setting

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/ramabmtr/heimdall/config"
)

func InitializeLogger() {
	logrus.SetOutput(os.Stdout)

	if config.AppMode == "production" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{})
	}

	if config.AppDebug {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
}
