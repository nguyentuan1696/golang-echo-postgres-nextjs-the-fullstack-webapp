package controller

import (
	"github.com/labstack/echo/v4"
	"thichlab-backend-docs/module/healthcheck/service"
)

type HealthCheckController struct {
	Service service.IHealthCheckService
}

func NewHealthCheckController(service *service.IHealthCheckService) *HealthCheckController {
	return &HealthCheckController{
		Service: service,
	}
}

func (ctl *HealthCheckController) GetStatus(c echo.Context) error {
	return nil
}
