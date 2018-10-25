package healthcheck

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ramabmtr/heimdall/helper/formatter"
)

func Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, formatter.BuildResponse("success", "pong", nil))
}
