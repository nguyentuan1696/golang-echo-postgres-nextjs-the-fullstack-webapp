package service

import (
	"context"
	"thichlab-backend-docs/dto"
	"thichlab-backend-docs/infrastructure/cache"
	"thichlab-backend-docs/module/account/repository"
)

type AccountService struct {
	AccountRepository repository.IAccountRepository
	Cache             cache.Client
}

func NewAccountService(repository repository.IAccountRepository, cache cache.Client) IAccountService {
	accountService := AccountService{
		Cache: cache,
	}
	accountService.AccountRepository = repository
	return &accountService
}

type IAccountService interface {
	AccountInsert(ctx context.Context, args *dto.AccountRegister) (*dto.AccountRegisterRes, error)
	AccountLogin(ctx context.Context, args *dto.AccountLoginReq) (string, string, error)
}
