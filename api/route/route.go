package route

import (
	"github.com/labstack/echo"
	"github.com/ramabmtr/heimdall/api/handler/healthcheck"
	"github.com/ramabmtr/heimdall/api/handler/verification"
)

func HealthCheckRouter(g *echo.Group) {
	g.GET(
		"/ping",
		healthcheck.Ping,
	)
}

func VerificationRouter(g *echo.Group) {
	g.POST(
		"/verification/send",
		verification.Send,
	)

	g.POST(
		"/verification/check",
		verification.Check,
	)
}
