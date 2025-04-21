package account

import (
	"go-api-starter/core/cache"
	"go-api-starter/core/database"
	"go-api-starter/core/middleware"
	"go-api-starter/modules/account/controller"
	"go-api-starter/modules/account/repository"
	"go-api-starter/modules/account/router"
	"go-api-starter/modules/account/service"

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
