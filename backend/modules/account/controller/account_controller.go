package controller

import (
	"thichlab-backend-slowpoke/core/utils"

	"github.com/labstack/echo/v4"
)


func (controller *AccountController) LockUser(c echo.Context) error {
	return nil
}

func (controller *AccountController) UnlockUser(c echo.Context) error {
	return nil
}

func (controller *AccountController) DeleteUser(c echo.Context) error {
	return nil
}

func (controller *AccountController) GetUsers(c echo.Context) error {
	ctx := c.Request().Context()

	pageNumber := utils.ToNumberWithDefault(c.QueryParam("pageNumber"), 1)
	pageSize := utils.ToNumberWithDefault(c.QueryParam("pageSize"), 20)

	resultGetUsers, err := controller.accountService.GetUsers(ctx, pageNumber, pageSize)
	if err != nil {
		return controller.NotFound("")
	}

	return controller.SuccessResponse(c, resultGetUsers, "Get users successfully")

}

func (controller *AccountController) GetUser(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Param("userId")
	resultGetUser, err := controller.accountService.GetUserById(ctx, userId)
	if err != nil {
		return controller.InternalServerError("Internal server error", err)
	}
	return controller.SuccessResponse(c, resultGetUser, "Get user success")
}

