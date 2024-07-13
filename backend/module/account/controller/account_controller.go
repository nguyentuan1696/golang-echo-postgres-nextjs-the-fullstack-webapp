package controller

import (
	"context"
	"github.com/labstack/echo/v4"
	"thichlab-backend-docs/dto"
	"thichlab-backend-docs/gerror"
	"thichlab-backend-docs/infrastructure/response"
	"thichlab-backend-docs/infrastructure/util"
)

func (controller *AccountController) CreateAccount(c echo.Context) error {
	var err error
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	p := new(dto.AccountRegister)

	err = c.Bind(&p)
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorBindData, err.Error(), util.FuncName())
		return controller.StatusBadRequestResponse(c, message, resp)
	}

	result, err := controller.AccountService.AccountInsert(ctx, p)
	if err != nil {
		message, resp := response.NewErrorResponse(gerror.ErrorSaveData, err.Error(), util.FuncName())
		return controller.StatusInternalServerErrorResponse(c, message, resp)
	}

	return controller.StatusOkResponse(c, result)
}
