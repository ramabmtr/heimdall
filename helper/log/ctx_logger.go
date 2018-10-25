package log

import (
	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/ramabmtr/heimdall/config"
)

func CreateContextLogger(ctx echo.Context, logger *logrus.Entry) {
	ctx.Set(config.Logger, logger)
}

func GetLogger(ctx echo.Context) *logrus.Entry {
	if ctx == nil {
		return logrus.NewEntry(logrus.StandardLogger())
	}

	logger := ctx.Get(config.Logger)

	if logger == nil {
		return logrus.NewEntry(logrus.StandardLogger())
	}

	return logger.(*logrus.Entry)
}
