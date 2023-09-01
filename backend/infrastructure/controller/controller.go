package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"thichlab-backend-docs/infrastructure/response"
)

type BaseController struct {
}

func (controller *BaseController) StatusOkResponse(c echo.Context, v any) error {

	return c.JSON(http.StatusOK, response.Response{
		Message: "success",
		Data:    v,
	})
}

func (controller *BaseController) StatusBadRequestResponse(c echo.Context, message string, err response.ErrorResponse) error {
	return controller.StatusErrorResponse(c, http.StatusBadRequest, message, err)
}

func (controller *BaseController) StatusInternalServerErrorResponse(c echo.Context, message string, err response.ErrorResponse) error {
	return controller.StatusErrorResponse(c, http.StatusInternalServerError, message, err)
}

func (controller *BaseController) StatusErrorResponse(c echo.Context, statusCode int, message string, err response.ErrorResponse) error {
	return c.JSON(statusCode, response.Response{
		Message: message,
		Data:    err,
	})
}
