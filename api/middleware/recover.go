package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"

	"github.com/labstack/echo"
	"github.com/ramabmtr/heimdall/helper/formatter"
	"github.com/ramabmtr/heimdall/helper/log"
)

func Recover() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			defer func() {
				if r := recover(); r != nil {
					err = errors.New(fmt.Sprintf("panic! %s", r))

					stack := make([]byte, 4<<10)
					length := runtime.Stack(stack, true)
					log.GetLogger(c).
						WithError(err).
						WithField("stack", string(stack[:length])).
						Error("recover from panic!")

					c.JSON(http.StatusInternalServerError, formatter.ErrorResponse("internal server error"))
				}
			}()
			return next(c)
		}
	}
}
