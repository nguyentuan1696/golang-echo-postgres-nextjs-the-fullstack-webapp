package account

import (
	"thichlab-backend-slowpoke/core/cache"
	"thichlab-backend-slowpoke/core/database"
	"thichlab-backend-slowpoke/core/middleware"
	"thichlab-backend-slowpoke/modules/account/controller"
	"thichlab-backend-slowpoke/modules/account/repository"
	"thichlab-backend-slowpoke/modules/account/router"
	"thichlab-backend-slowpoke/modules/account/service"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, db database.Database, cache *cache.Cache) {
	repository := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(repository, cache)
	middlewares := middleware.NewMiddleware(accountService)

	// Update: pass only the controller
	router.NewAccountRouter(
		controller.NewAccountController(accountService),
	).Setup(e, middlewares) // Pass middleware to Setup instead
}
