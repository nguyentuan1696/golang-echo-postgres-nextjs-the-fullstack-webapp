package healthcheck

import (
	"github.com/labstack/echo/v4"
	"thichlab-backend-docs/module/healthcheck/controller"
	"thichlab-backend-docs/module/healthcheck/service"
)

var mHealthCheckController *controller.HealthCheckController

func Initialize(e *echo.Echo) {
	healthCheckService := service.NewHealthCheckService()
	mHealthCheckController = controller.NewHealthCheckController(&healthCheckService)

	initRouter(e)
}

func initRouter(e *echo.Echo) {
	e.GET("/docs/v2/healthcheck/status", mHealthCheckController.GetStatus)
}
