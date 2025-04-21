package repository

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"thichlab-backend-slowpoke/core/logger"
	"thichlab-backend-slowpoke/modules/account/entity"

	"github.com/google/uuid"
)

func (r *AccountRepository) GetUserByEmailOrUserNameOrId(ctx context.Context, email, userName string, userId uuid.UUID) (*entity.User, error) {
	var (
		queryBuilder strings.Builder
		args         []any
	)

	queryBuilder.WriteString(`SELECT * FROM users WHERE `)

	switch {
	case email != "" && userName != "":
		queryBuilder.WriteString(`email = $1 OR user_name = $2 LIMIT 1`)
		args = append(args, email, userName)
	case email != "":
		queryBuilder.WriteString(`email = $1 LIMIT 1`)
		args = append(args, email)
	case userName != "":
		queryBuilder.WriteString(`user_name = $1 LIMIT 1`)
		args = append(args, userName)
	case userId != uuid.Nil:
		queryBuilder.WriteString(`id = $1 LIMIT 1`)
		args = append(args, userId)
	default:
		return nil, nil
	}

	var user entity.User
	err := r.DB.GetContext(ctx, &user, queryBuilder.String(), args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *AccountRepository) CreateAccount(ctx context.Context, user *entity.User) (*entity.User, error) {

	query := `
		INSERT INTO users (user_name, email, password)
		VALUES (:user_name, :email, :password)
		RETURNING id, user_name, email
	`

	rows, err := r.DB.NamedQueryContext(ctx, query, user)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.StructScan(user)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (r *AccountRepository) UpdatePassword(ctx context.Context, user *entity.User) error {
	query := `
        UPDATE users 
        SET password = :password
        WHERE id = :id
    `

	result, err := r.DB.NamedExecContext(ctx, query, user)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("no user found to update")
	}

	return nil
}

func (r *AccountRepository) GetUsers(ctx context.Context, pageNumber, pageSize int) (*entity.PaginatedUsers, error) {
	offset := (pageNumber - 1) * pageSize

	// Get total count
	var totalItems int
	countQuery := `SELECT COUNT(*) FROM users table`
	err := r.DB.GetContext(ctx, &totalItems, countQuery)
	if err != nil {
		logger.Error("AccountRepository:GetUsers:Error when count users", "error", err)
		return nil, err
	}

	// Get paginated users
	query := `
        SELECT id, user_name, email, created_at, updated_at 
        FROM users 
        ORDER BY created_at DESC
        LIMIT $1 OFFSET $2
    `

	users := []*entity.User{}
	err = r.DB.SelectContext(ctx, &users, query, pageSize, offset)
	if err != nil {
		logger.Error("AccountRepository:GetUsers:Error when get users", "error", err)
		return nil, err
	}

	totalPages := (totalItems + pageSize - 1) / pageSize

	return &entity.PaginatedUsers{
		Items:       users,
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		CurrentPage: pageNumber,
		PageSize:    pageSize,
	}, nil
}


