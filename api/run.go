package api

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	appMiddleware "github.com/ramabmtr/heimdall/api/middleware"
	"github.com/ramabmtr/heimdall/api/route"
	"github.com/ramabmtr/heimdall/config"
	"github.com/ramabmtr/heimdall/helper/formatter"
	"github.com/ramabmtr/heimdall/helper/log"
	"github.com/ramabmtr/heimdall/helper/message"
	"github.com/satori/go.uuid"
)

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func Run() {
	e := echo.New()

	echo.NotFoundHandler = func(c echo.Context) error {
		return c.JSON(http.StatusNotFound, formatter.FailResponse("no route found"))
	}

	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: func() string {
			return uuid.NewV4().String()
		},
	}))
	e.Use(middleware.Secure())
	e.Use(appMiddleware.RequestLogger())
	e.Use(appMiddleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	g := e.Group("v1")
	{
		route.HealthCheckRouter(g)
		route.VerificationRouter(g)
	}

	e.HideBanner = true

	message.ServiceNameMessage()
	message.EnvInfoMessage()
	err := e.Start(config.AppAddress)
	log.GetLogger(nil).WithError(err).Fatal("fail to initialize application")
}
