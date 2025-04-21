package controller

import (
	"go-api-starter/core/controller"
	"go-api-starter/modules/account/service"
)

type AccountController struct {
	controller.BaseController
	accountService service.IAccountService
}

func NewAccountController(service service.IAccountService) *AccountController {

	return &AccountController{
		BaseController: controller.NewBaseController(),
		accountService: service,
	}
}
