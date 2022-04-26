package echoHealth

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func SetupHealthChecks(config HealthCheckConfig, healthChecker []HealthChecker, e *echo.Echo, middleware echo.MiddlewareFunc) {
	e.GET(config.ApplicationName+"/_health", func(c echo.Context) error {
		messages := make([]string, len(healthChecker))
		statusCode := http.StatusOK
		isErr := false
		for i, checker := range healthChecker {
			messages[i], isErr = checker.CheckHealth()
			if isErr {
				statusCode = http.StatusBadRequest
			}
		}

		return c.JSON(statusCode, Response{
			messages: messages,
		})
	}, middleware)
}
