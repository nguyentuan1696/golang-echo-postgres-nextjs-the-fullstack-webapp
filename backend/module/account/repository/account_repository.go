package repository

import (
	"context"
	"fmt"
	"thichlab-backend-docs/constant"
	"thichlab-backend-docs/dto"
)

func (repository *AccountRepository) DBAccountInsert(ctx context.Context, args *dto.AccountRegister) error {
	query := fmt.Sprintf("INSERT INTO %s (id, username, email, password, created_at, updated_at) VALUES ('%s', '%s', '%s', '%s',  %d, %d)",
		constant.TableAccount, args.Id, args.Username, args.Email, args.Password, args.CreatedAt, args.UpdatedAt,
	)

	_, err := repository.Postgres.SQLxDBContext.ExecContext(ctx, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return err
	}

	return nil
}

func (repository *AccountRepository) GetInfoUserByUsername(ctx context.Context, username string) (*dto.AccountRegister, error) {
	r := new(dto.AccountRegister)
	query := fmt.Sprintf("select id, username, email, password from accounts where username = '%s'", username)
	err := repository.Postgres.SQLxDBContext.GetContext(ctx, r, query)
	if err != nil {
		repository.Postgres.HandleError(err, query)
		return nil, err
	}
	return r, nil
}
