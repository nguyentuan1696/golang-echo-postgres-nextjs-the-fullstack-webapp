package controller

import (
	"thichlab-backend-docs/infrastructure/controller"
	"thichlab-backend-docs/module/docs/service"
)

type DocsController struct {
	controller.BaseController
	DocsService service.IDocsService
}
