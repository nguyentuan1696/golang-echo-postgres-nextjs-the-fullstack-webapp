package controller

import (
	"thichlab-backend-slowpoke/core/controller"
	"thichlab-backend-slowpoke/modules/account/service"
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
