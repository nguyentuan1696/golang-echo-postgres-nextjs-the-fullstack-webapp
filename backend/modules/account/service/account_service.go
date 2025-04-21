package service

import (
	"context"
	"thichlab-backend-slowpoke/core/errors"

	"thichlab-backend-slowpoke/core/logger"
	"thichlab-backend-slowpoke/core/utils"
	"thichlab-backend-slowpoke/modules/account/dto"
	"thichlab-backend-slowpoke/modules/account/mapper"
	"time"
)

func (s *AccountService) GetUserById(ctx context.Context, userId string) (*dto.UserResponse, *errors.AppError) {
	return nil, nil

}

func (s *AccountService) GetUsers(ctx context.Context, pageNumber, pageSize int) (*dto.PaginatedUsersResponse, *errors.AppError) {

	ctx, cancel := utils.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	resultGetUsers, err := s.repo.GetUsers(ctx, pageNumber, pageSize)
	if err != nil {
		logger.Error("AccountService:GetUsers:Failed to get users", "error", err)
		return nil, errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}
	// Convert to DTO
	usersDTO := mapper.ToPaginatedUsersResponse(resultGetUsers)
	return usersDTO, nil
}
