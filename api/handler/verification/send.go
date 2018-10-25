package verification

import (
	"github.com/ramabmtr/heimdall/repository/database/memcached"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/ramabmtr/heimdall/config"
	"github.com/ramabmtr/heimdall/helper/formatter"
	"github.com/ramabmtr/heimdall/model"
	"github.com/ramabmtr/heimdall/repository/database/redis"
	"github.com/ramabmtr/heimdall/repository/external/nexmo"
	"github.com/ramabmtr/heimdall/repository/external/postmark"
	"github.com/ramabmtr/heimdall/repository/external/twilio"
	"github.com/ramabmtr/heimdall/service"
)

type (
	sendVerificationParams struct {
		Service string      `json:"service"`
		SendTo  interface{} `json:"send_to" validate:"required"`
	}
)

func Send(c echo.Context) error {
	params := new(sendVerificationParams)
	if err := c.Bind(params); err != nil {
		return c.JSON(http.StatusBadRequest, formatter.FailResponse(err.Error()))
	}
	if err := c.Validate(params); err != nil {
		return c.JSON(http.StatusBadRequest, formatter.FailResponse(err.Error()))
	}

	verificationDB := GetVerificationDB(c)
	verificationRepo := GetVerificationRepo(c, params.Service, verificationDB)

	verifyService := service.NewVerificationService(c, verificationRepo)

	if err := verifyService.SendVerificationCode(params.SendTo); err != nil {
		status := http.StatusInternalServerError
		if _, ok := err.(validator.ValidationErrors); ok {
			status = http.StatusBadRequest
		}
		return c.JSON(status, formatter.ErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, formatter.SuccessResponse())
}

func GetVerificationDB(ctx echo.Context) (verificationDB model.VerificationDatabaseRepository) {
	switch config.DatabaseType {
	case config.DBTypeMemCached:
		verificationDB = memcached.NewVerificationDatabaseRepository(ctx)
	case config.DBTypeRedis:
		fallthrough
	default:
		verificationDB = redis.NewVerificationDatabaseRepository(ctx)
	}

	return
}

func GetVerificationRepo(
	ctx echo.Context,
	serviceName string,
	vDB model.VerificationDatabaseRepository,
) (verificationRepo model.VerificationRepository) {
	switch serviceName {
	case config.ServiceNamePostmark:
		verificationRepo = postmark.NewVerificationRepository(ctx, vDB)
	case config.ServiceNameNexmo:
		verificationRepo = nexmo.NewVerificationRepository(ctx, vDB)
	case config.ServiceNameTwilio:
		fallthrough
	default:
		verificationRepo = twilio.NewVerificationRepository(ctx, vDB)
	}

	return
}
