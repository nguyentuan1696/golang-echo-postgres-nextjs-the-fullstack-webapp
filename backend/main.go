package main

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"thichlab-backend-docs/constant"
	"thichlab-backend-docs/infrastructure/cache"
	"thichlab-backend-docs/infrastructure/logger"
	"thichlab-backend-docs/infrastructure/repository"
	"thichlab-backend-docs/infrastructure/search"
	"thichlab-backend-docs/module/account"
	"thichlab-backend-docs/module/docs"
	"thichlab-backend-docs/module/healthcheck"
)

func init() {
	viper.SetConfigFile(`config.json`)

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	if viper.GetBool(`Debug`) {
		fmt.Println("service RUN on DEBUG mode")
	} else {
		fmt.Println("Service RUN on PRODUCTION mode")
	}
}

func main() {

	//------------------------------------
	// CONFIGURE POSTGRES DB
	//------------------------------------
	logPath := viper.GetString("Log.Path")
	logPrefix := viper.GetString("Log.Prefix")
	logger.NewLogger(logPath, logPrefix)

	//------------------------------------
	// CONFIGURE POSTGRES DB
	//------------------------------------

	postgresHost := viper.GetString("Postgres.Host")
	postgresPort := viper.GetInt("Postgres.Port")
	postgresUserName := viper.GetString("Postgres.UserName")
	postgresPassword := viper.GetString("Postgres.Password")
	postgresDB := viper.GetString("Postgres.Database")
	postgresUri := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=verify-full",
		postgresHost, postgresPort, postgresUserName, postgresPassword, postgresDB)
	sqlDBContext, sqlxDBContext := repository.InitializeConnection(postgresUri)
	defer func(sqlDBContext *sql.DB) {
		err := sqlDBContext.Close()
		if err != nil {

		}
	}(sqlDBContext)
	defer func(sqlxDBContext *sqlx.DB) {
		err := sqlxDBContext.Close()
		if err != nil {

		}
	}(sqlxDBContext)
	fmt.Println("Successfully connected to postgres")

	//------------------------------------
	// CONFIGURE Redis:Cache
	//------------------------------------

	hostCache := viper.GetString("Redis.Host")
	passWord := viper.GetString("Redis.Password")
	poolSize := viper.GetInt("Redis.PoolSize")
	minIdleConns := viper.GetInt("Redis.MinIdleConns")
	db := viper.GetInt("Redis.DB")
	cacheClient := cache.Client{}
	cacheClient.InitializeConnection(hostCache, passWord, poolSize, minIdleConns, db)
	fmt.Println("Successfully connected to redis cache")

	//------------------------------------
	// CONFIGURE Redis:Search
	//------------------------------------
	hostSearch := viper.GetString("RedisSearch.Host")
	name := viper.GetString("RedisSearch.IndexName")
	password := viper.GetString("RedisSearch.Password")
	searchClient := search.Client{}
	searchClient.InitializeConnection(hostSearch, name, password)
	fmt.Println("Successfully connected to redis search")

	//------------------------------------
	// CONFIGURE ECHO
	//------------------------------------
	e := echo.New()

	e.Server.SetKeepAlivesEnabled(false)
	e.Server.ReadTimeout = constant.TimeoutServerDefault
	e.Server.WriteTimeout = constant.TimeoutServerDefault

	e.Use(middleware.CORS())

	//------------------------------------
	// INITIALIZE MODULES
	//------------------------------------

	// Init HealthCheck module
	healthcheck.Initialize(e)

	// Init docs module
	docs.Initialize(e, sqlDBContext, sqlxDBContext, cacheClient, searchClient)

	// Init account module
	account.Initialize(e, sqlDBContext, sqlxDBContext, cacheClient)

	//------------------------------------
	// START APPLICATION
	//------------------------------------
	err := e.Start(viper.GetString("Server.Address"))
	if err != nil {
		logger.Error("ServerStart: - ERROR: %s", err)
		return
	}
}
