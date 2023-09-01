package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"thichlab-backend-docs/infrastructure/logger"
	"thichlab-backend-docs/infrastructure/util"
	"time"
)

type PostgresRepository struct {
	DbContext     *sql.DB
	SQLxDBContext *sqlx.DB
}

func (repository *PostgresRepository) SetDbContext(dbContext *sql.DB, sqlxDBContext *sqlx.DB) {
	repository.DbContext = dbContext
	repository.SQLxDBContext = sqlxDBContext
}

func (repository *PostgresRepository) HandleError(err error, query string) {
	logger.Error("[POSTGRES] - Error: %s, Query: %s", err, query)
}

func InitializeConnection(postgresUri string) (*sql.DB, *sqlx.DB) {
	maxOpenConnections := viper.GetString("Postgres.MaxOpenConnections")
	maxIdleConnections := viper.GetString("Postgres.MaxIdleConnections")
	connMaxLifetime := viper.GetString("Postgres.ConnMaxLifetime")

	maxOpenConnectionsInt := util.ToInt(maxOpenConnections)
	maxIdleConnectionsInt := util.ToInt(maxIdleConnections)
	connMaxLifetimeInt := util.ToInt(connMaxLifetime)
	connMaxLifetimeHours := time.Duration(connMaxLifetimeInt) * time.Hour

	// SQL
	sqlDBContext, err := sql.Open("postgres", postgresUri)
	if err != nil {
		logger.Error("[Postgres] - Init connection exception: %s", err)
		panic(err)
	}
	sqlDBContext.SetMaxOpenConns(maxOpenConnectionsInt)
	sqlDBContext.SetMaxIdleConns(maxIdleConnectionsInt)
	sqlDBContext.SetConnMaxLifetime(connMaxLifetimeHours)

	err = sqlDBContext.Ping()
	if err != nil {
		logger.Error("[Postgres] - Init connection exception: %s", err)
		panic(err)
	}
	// SQLx
	sqlxDBContext, err := sqlx.Open("postgres", postgresUri)
	if err != nil {
		logger.Error("[Postgres] - Init connection exception: %s", err)
		panic(err)
	}

	err = sqlxDBContext.Ping()
	if err != nil {
		logger.Error("[Postgres] - Init connection exception: %s", err)
		panic(err)
	}
	sqlxDBContext.SetMaxOpenConns(maxOpenConnectionsInt)
	sqlxDBContext.SetMaxIdleConns(maxIdleConnectionsInt)
	sqlxDBContext.SetConnMaxLifetime(connMaxLifetimeHours)

	return sqlDBContext, sqlxDBContext
}
