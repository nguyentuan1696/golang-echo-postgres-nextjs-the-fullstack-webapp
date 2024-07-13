package account

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"thichlab-backend-docs/infrastructure/cache"
	"thichlab-backend-docs/module/account/controller"
	"thichlab-backend-docs/module/account/repository"
	"thichlab-backend-docs/module/account/service"
)

var mAccountController *controller.AccountController

func Initialize(e *echo.Echo, dbContext *sql.DB, sqlxDBContext *sqlx.DB, cache cache.Client) {

	AccountRepository := repository.NewAccountRepository(dbContext, sqlxDBContext)
	AccountService := service.NewAccountService(AccountRepository, cache)

	mAccountController = &controller.AccountController{
		AccountService: AccountService,
	}

	// For doing something in feature
	//ticker := time.NewTicker(time.Second)
	//go func() {
	//	for range ticker.C {
	//		fmt.Println("Tick")
	//	}
	//}()

	initRoute(e)

}

func initRoute(e *echo.Echo) {
	route := e.Group("/account/api/v1/")
	route.POST("signup", mAccountController.CreateAccount)
	route.POST("login", mAccountController.LoginAccount)

}
