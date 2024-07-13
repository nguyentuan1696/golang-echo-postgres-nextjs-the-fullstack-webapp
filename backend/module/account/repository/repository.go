package repository

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"thichlab-backend-docs/dto"
	"thichlab-backend-docs/infrastructure/repository"
)

type AccountRepository struct {
	Postgres repository.PostgresRepository
}

func NewAccountRepository(dbContext *sql.DB, sqlxDBContext *sqlx.DB) IAccountRepository {
	accountRepository := AccountRepository{}
	accountRepository.Postgres.SetDbContext(dbContext, sqlxDBContext)
	return &accountRepository
}

type IAccountRepository interface {
	DBAccountInsert(ctx context.Context, args *dto.AccountRegister) error
	GetInfoUserByUsername(ctx context.Context, username string) (*dto.AccountRegister, error)
}
