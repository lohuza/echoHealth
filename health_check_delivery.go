package echoHealth

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func SetupHealthChecks(config HealthCheckConfig, healthChecker []HealthChecker, e *echo.Echo, middleware echo.MiddlewareFunc) {
	e.GET(config.ApplicationName+"/_health", func(c echo.Context) error {
		messages := make([]string, len(healthChecker))
		for i, checker := range healthChecker {
			messages[i] = checker.CheckHealth()
		}

		return c.JSON(http.StatusOK, messages)
	}, middleware)
}
