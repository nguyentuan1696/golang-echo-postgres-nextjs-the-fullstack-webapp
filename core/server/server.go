package server

import (
	"context"
	"flag"
	"fmt"
	"go-api-starter/core/cache"
	"go-api-starter/core/config"
	"go-api-starter/core/database"
	"go-api-starter/core/logger"
	"go-api-starter/core/middleware"
	"go-api-starter/modules/account"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

type Server struct {
	echo  *echo.Echo
	addr  string
	cache *cache.Cache
	db    database.Database
}

func initEnvironment() (config.Environment, error) {
	env := flag.String("env", "dev", "Environment (dev/prod)")
	flag.Parse()

	switch *env {
	case "dev":
		return config.DevEnvironment, nil
	case "prod":
		return config.ProdEnvironment, nil
	default:
		return "", fmt.Errorf("invalid environment. Use 'dev' or 'prod'")
	}
}

func initServer() (*Server, error) {
	environment, err := initEnvironment()
	if err != nil {
		return nil, err
	}

	if errInitConfig := config.Init(environment); errInitConfig != nil {
		return nil, fmt.Errorf("failed to initialize config: %w", errInitConfig)
	}

	cfg := config.Get()

	// Initialize logger
	if errInitLogger := logger.Init(logger.LogConfig{
		Level:    logger.LogLevelDebug,
		FilePath: "app.log",
	}); errInitLogger != nil {
		return nil, fmt.Errorf("failed to initialize logger: %w", errInitLogger)
	}

	// Initialize database
	db, err := database.InitDB(database.DatabaseConfig{
		Host:                   cfg.Database.Host,
		Port:                   cfg.Database.Port,
		User:                   cfg.Database.User,
		Password:               cfg.Database.Password,
		DBName:                 cfg.Database.DBName,
		MaxOpenConns:           10, // Default value
		MaxIdleConns:           5,  // Default value
		ConnMaxLifetime:        60, // Default value in minutes
		SSLMode:                "disable",
		ConnectTimeout:         10,
		StatementTimeout:       30,
		IdleInTxSessionTimeout: 60,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	logger.Info("Server initializing",
		"environment", environment,
		"host", cfg.Server.Host,
		"port", cfg.Server.Port,
	)

	// Initialize Redis cache
	redisCache := cache.NewCache(
		cfg.Redis.Address,
		cfg.Redis.Password,
		cfg.Redis.DB,
	)

	e := echo.New()

	// Middleware
	e.Use(middleware.LoggerMiddleware())
	e.Use(middleware.CORSMiddleware())

	// Initialize modules
	account.Init(e, db, redisCache)

	return &Server{
		echo:  e,
		addr:  fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
		cache: redisCache,
		db:    db,
	}, nil
}

func (s *Server) start() error {
	logger.Info("Starting HTTP server", "address", s.addr)

	go func() {
		if err := s.echo.Start(s.addr); err != nil {
			logger.Info("Shutting down server", "error", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.echo.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown server gracefully: %w", err)
	}

	// Close Redis connection
	if err := s.cache.Close(); err != nil {
		logger.Error("Failed to close Redis connection", "error", err)
	}

	logger.Info("Server shutdown complete")
	return nil
}

func Run() error {
	srv, err := initServer()
	if err != nil {
		return err
	}
	return srv.start()
}
