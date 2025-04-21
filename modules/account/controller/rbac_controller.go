package controller

import (
	"thichlab-backend-slowpoke/modules/account/dto"
	"thichlab-backend-slowpoke/modules/account/validator"

	"github.com/labstack/echo/v4"
)

func (controller *AccountController) GetRoles(c echo.Context) error {
	ctx := c.Request().Context()
	resultGetRoles, err := controller.accountService.GetRoles(ctx)
	if err != nil {
		return controller.InternalServerError("Internal server error", err)
	}
	return controller.SuccessResponse(c, resultGetRoles, "Get roles success")
}

func (controller *AccountController) CreateRole(c echo.Context) error {
	ctx := c.Request().Context()

	requestData := new(dto.CreateRoleRequest)
	err := c.Bind(requestData)
	if err != nil {
		return controller.BadRequest("Invalid request data", err)
	}

	resultValidation := validator.ValidateCreateRole(requestData)
	if !resultValidation.Valid {
		return controller.BadRequest("Invalid request data", resultValidation.Errors)
	}

	resultCreateRole := controller.accountService.CreateRole(ctx, requestData)

	if resultCreateRole != nil {
		return controller.InternalServerError("Internal server error", resultCreateRole)
	}

	return controller.SuccessResponse(c, nil, "Create role success")

}

func (controller *AccountController) CreatePermission(c echo.Context) error {
	ctx := c.Request().Context()
	requestData := new(dto.CreatePermissionRequest)
	err := c.Bind(requestData)
	if err != nil {
		return controller.BadRequest("Invalid request data", err)
	}
	resultValidation := validator.ValidateCreatePermission(requestData)
	if !resultValidation.Valid {
		return controller.BadRequest("Invalid request data", resultValidation.Errors)
	}
	resultCreatePermission := controller.accountService.CreatePermission(ctx, requestData)
	if resultCreatePermission != nil {
		return controller.InternalServerError("Internal server error", resultCreatePermission)
	}
	return controller.SuccessResponse(c, nil, "Create permission success")
}

func (controller *AccountController) GetPermissions(c echo.Context) error {
	ctx := c.Request().Context()
	resultGetPermissions, err := controller.accountService.GetPermissions(ctx)
	if err != nil {
		return controller.InternalServerError("Internal server error", err)
	}
	return controller.SuccessResponse(c, resultGetPermissions, "Get permissions success")
}

func (controller *AccountController) AssignPermissionToRole(c echo.Context) error {
	ctx := c.Request().Context()
	requestData := new(dto.AssignPermissionToRoleRequest)
	err := c.Bind(requestData)
	if err != nil {
		return controller.BadRequest("Invalid request data", err)
	}
	resultValidation := validator.ValidateAssignPermissionToRole(requestData)
	if !resultValidation.Valid {
		return controller.BadRequest("Invalid request data", resultValidation.Errors)
	}
	resultAssignPermissionToRole := controller.accountService.AssignPermissionToRole(ctx, requestData.RoleId, requestData.PermissionId)
	if resultAssignPermissionToRole != nil {
		return controller.InternalServerError("Internal server error", resultAssignPermissionToRole)
	}
	return controller.SuccessResponse(c, nil, "Assign permission to role success")
}

func (controller *AccountController) AssignRoleToUser(c echo.Context) error {
	ctx := c.Request().Context()
	requestData := new(dto.AssignRoleToUserRequest)
	err := c.Bind(requestData)
	if err != nil {
		return controller.BadRequest("Invalid request data", err)
	}
	resultValidation := validator.ValidateAssignRoleToUser(requestData)
	if !resultValidation.Valid {
		return controller.BadRequest("Invalid request data", resultValidation.Errors)
	}
	resultAssignRoleToUser := controller.accountService.AssignRoleToUser(ctx, requestData.UserId, requestData.RoleId)
	if resultAssignRoleToUser != nil {
		return controller.InternalServerError("Internal server error", resultAssignRoleToUser)
	}
	return controller.SuccessResponse(c, nil, "Assign role to user success")
}
