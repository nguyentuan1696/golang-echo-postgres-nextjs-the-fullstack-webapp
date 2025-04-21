package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Response types
type (
	SuccessResponse struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Data    any    `json:"data,omitempty"`
	}

	ErrorResponse struct {
		Status  string `json:"status"`
		Code    int    `json:"code"`
		Message string `json:"message"`
		Details any    `json:"details,omitempty"`
	}

	ValidationError struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	}

	ValidationResponse struct {
		Success bool              `json:"success"`
		Message string            `json:"message"`
		Errors  []ValidationError `json:"errors"`
	}
)

// Response handler interface and implementation
type BaseController interface {
	BadRequest(message string, details ...any) *echo.HTTPError
	InternalServerError(message string, details ...any) *echo.HTTPError
	NotFound(message string, details ...any) *echo.HTTPError
	Unauthorized(message string, details ...any) *echo.HTTPError
	Forbidden(message string, details ...any) *echo.HTTPError
	SuccessResponse(c echo.Context, data any, message string) error
	ErrorResponse(c echo.Context, err error) error
}

type responseHandler struct{}

func NewBaseController() BaseController {
	return &responseHandler{}
}

// Success response functions
func NewSuccessResponse(data any, message string) *SuccessResponse {
	return &SuccessResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	}
}

// Error response functions
func NewErrorResponse(code int, message string, details ...interface{}) *echo.HTTPError {
	err := &ErrorResponse{
		Code:    code,
		Message: message,
	}
	if len(details) > 0 {
		err.Details = details[0]
	}
	return echo.NewHTTPError(code, err)
}

// Validation functions
func NewValidationError(field, message string) ValidationError {
	return ValidationError{
		Field:   field,
		Message: message,
	}
}

// HTTP Error handlers
func (h *responseHandler) BadRequest(message string, details ...interface{}) *echo.HTTPError {
	return NewErrorResponse(http.StatusBadRequest, message, details...)
}

func (h *responseHandler) InternalServerError(message string, details ...interface{}) *echo.HTTPError {
	return NewErrorResponse(http.StatusInternalServerError, message, details...)
}

func (h *responseHandler) NotFound(message string, details ...interface{}) *echo.HTTPError {
	return NewErrorResponse(http.StatusNotFound, message, details...)
}

func (h *responseHandler) Unauthorized(message string, details ...interface{}) *echo.HTTPError {
	return NewErrorResponse(http.StatusUnauthorized, message, details...)
}

func (h *responseHandler) Forbidden(message string, details ...interface{}) *echo.HTTPError {
	return NewErrorResponse(http.StatusForbidden, message, details...)
}

func (h *responseHandler) ValidationError(message string, details ...interface{}) *echo.HTTPError {
	return NewErrorResponse(http.StatusBadRequest, message, details...)
}

func (h *responseHandler) SuccessResponse(c echo.Context, data any, message string) error {
	return c.JSON(http.StatusOK, NewSuccessResponse(data, message))
}

func (h *responseHandler) ErrorResponse(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, NewErrorResponse(http.StatusInternalServerError, err.Error()))
}
