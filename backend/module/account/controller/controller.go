package controller

import (
	"thichlab-backend-docs/infrastructure/controller"
	"thichlab-backend-docs/module/account/service"
)

type AccountController struct {
	controller.BaseController
	AccountService service.IAccountService
}
