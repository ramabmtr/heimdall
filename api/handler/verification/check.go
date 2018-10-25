package verification

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ramabmtr/heimdall/helper/formatter"
	"github.com/ramabmtr/heimdall/service"
)

type (
	checkVerificationParams struct {
		Service  string      `json:"service"`
		Code     string      `json:"code" validate:"required"`
		CheckKey interface{} `json:"check_key" validate:"required"`
	}
)

func Check(c echo.Context) error {
	params := new(checkVerificationParams)
	if err := c.Bind(params); err != nil {
		return c.JSON(http.StatusBadRequest, formatter.FailResponse(err.Error()))
	}
	if err := c.Validate(params); err != nil {
		return c.JSON(http.StatusBadRequest, formatter.FailResponse(err.Error()))
	}

	verificationDB := GetVerificationDB(c)
	verificationRepo := GetVerificationRepo(c, params.Service, verificationDB)

	verifyService := service.NewVerificationService(c, verificationRepo)

	pass, err := verifyService.CheckVerificationCode(params.CheckKey, params.Code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, formatter.ErrorResponse(err.Error()))
	}

	if !pass {
		return c.JSON(http.StatusPreconditionFailed, formatter.ErrorResponse("code not match"))
	}

	return c.JSON(http.StatusOK, formatter.SuccessResponse())
}
