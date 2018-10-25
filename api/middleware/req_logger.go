package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/labstack/echo"
	"github.com/ramabmtr/heimdall/config"
	"github.com/ramabmtr/heimdall/helper/log"
)

func RequestLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			req := c.Request()
			res := c.Response()

			l := log.GetLogger(c).WithField("req_id", res.Header().Get(echo.HeaderXRequestID)).
				WithField("user_agent", req.Header.Get("User-Agent"))
			log.CreateContextLogger(c, l)

			log.GetLogger(c).
				Info(fmt.Sprintf("request started: %s %s", req.Method, req.URL))

			if config.AppDebug {
				var bodyBytes []byte

				if c.Request().Body != nil {
					bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
				}

				c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

				var body interface{}

				if err = json.Unmarshal(bodyBytes, &body); err != nil {
					return err
				}

				bodyBytes, _ = json.Marshal(body)

				log.GetLogger(c).
					WithField("headers", req.Header).
					WithField("req_body", string(bodyBytes)).
					Debug("request detail")
			}

			err = next(c)
			if err != nil {
				log.GetLogger(c).WithError(err).Error("request error")
			}

			finishLog := log.GetLogger(c).WithField("status", res.Status)
			finishLogMsg := fmt.Sprintf("request finished: %s %s", req.Method, req.URL)

			switch {
			case res.Status > 499:
				finishLog.Error(finishLogMsg)
			case res.Status > 399:
				finishLog.Warn(finishLogMsg)
			default:
				finishLog.Info(finishLogMsg)
			}

			return err
		}
	}
}
