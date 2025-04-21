package database

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"thichlab-backend-slowpoke/core/logger"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type IDatabase interface {
	ExecContext(ctx context.Context, query string, args ...any) error
	GetContext(ctx context.Context, dest interface{}, query string, args ...any) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...any) error
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	NamedQueryContext(ctx context.Context, query string, arg any) (*sqlx.Rows, error)
	NamedExecContext(ctx context.Context, query string, arg any) (sql.Result, error)
}

type Database struct {
	db   *sql.DB
	sqlx *sqlx.DB
}

type DatabaseConfig struct {
	Host                   string
	Port                   int
	User                   string
	Password               string
	DBName                 string
	MaxOpenConns           int
	MaxIdleConns           int
	ConnMaxLifetime        int    // in minutes
	SSLMode                string // disable, require, verify-ca, verify-full
	ConnectTimeout         int    // in seconds
	StatementTimeout       int    // in seconds
	IdleInTxSessionTimeout int    // in seconds
}

var (
	instance *Database
	once     sync.Once
)

func GetDB() IDatabase {
	return instance
}

func InitDB(config DatabaseConfig) (Database, error) {
	logger.Info("Initializing database...")
	var db Database
	var err error

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName)

	sqlxDB, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		logger.Error("Failed to connect to database", "error", err)
		return Database{}, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB := sqlxDB.DB
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(config.ConnMaxLifetime) * time.Minute)

	if err = sqlDB.Ping(); err != nil {
		logger.Error("Failed to ping database", "error", err)
		return Database{}, fmt.Errorf("failed to ping database: %w", err)
	}

	db = Database{
		db:   sqlDB,
		sqlx: sqlxDB,
	}

	logger.Info("Database initialized successfully",
		"maxOpenConns", config.MaxOpenConns,
		"maxIdleConns", config.MaxIdleConns,
		"connMaxLifetime", config.ConnMaxLifetime,
	)

	return db, nil
}

func (d *Database) ExecContext(ctx context.Context, query string, args ...any) error {
	_, err := d.sqlx.ExecContext(ctx, query, args...)
	return err
}

func (d *Database) GetContext(ctx context.Context, dest any, query string, args ...any) error {
	return d.sqlx.GetContext(ctx, dest, query, args...)
}

func (d *Database) SelectContext(ctx context.Context, dest any, query string, args ...any) error {
	return d.sqlx.SelectContext(ctx, dest, query, args...)
}

func (d *Database) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	return d.db.QueryRowContext(ctx, query, args...)
}

func (d *Database) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return d.db.QueryContext(ctx, query, args...)
}

func (d *Database) NamedQueryContext(ctx context.Context, query string, arg any) (*sqlx.Rows, error) {
	return d.sqlx.NamedQueryContext(ctx, query, arg)
}

func (d *Database) NamedExecContext(ctx context.Context, query string, arg any) (sql.Result, error) {
	return d.sqlx.NamedExecContext(ctx, query, arg)
}


