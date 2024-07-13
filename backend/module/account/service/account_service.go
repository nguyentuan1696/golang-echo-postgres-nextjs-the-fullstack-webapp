package service

import (
	"context"
	"github.com/google/uuid"
	"strings"
	"thichlab-backend-docs/dto"
	"thichlab-backend-docs/infrastructure/logger"
	"thichlab-backend-docs/infrastructure/util"
)

func (service *AccountService) AccountInsert(ctx context.Context, args *dto.AccountRegister) (*dto.AccountRegisterRes, error) {

	args.Username = strings.TrimSpace(args.Username)
	args.Username = strings.ToLower(args.Username)
	args.Email = strings.ToLower(args.Email)

	AccountRegister := new(dto.AccountRegisterRes)

	err := util.PasswordValidator(args.Password)
	if err != nil {
		return AccountRegister, nil
	}

	err = util.UsernameValidator(args.Username)
	if err != nil {
		return AccountRegister, nil
	}

	err = util.EmailValidator(args.Email)
	if err != nil {
		return AccountRegister, nil
	}

	hashPassword, _ := util.HashPassword(args.Password)

	accountDto := dto.AccountRegister{
		Id:        uuid.New(),
		Username:  args.Username,
		Email:     args.Email,
		Password:  hashPassword,
		CreatedAt: util.NowUnixTimeMillisecond(),
		UpdatedAt: util.NowUnixTimeMillisecond(),
	}

	dataRes := dto.AccountRegisterRes{
		UserId:   accountDto.Id,
		Username: accountDto.Username,
	}

	err = service.AccountRepository.DBAccountInsert(ctx, &accountDto)
	if err != nil {
		logger.Error("AccountService:AccountInsert - ERROR: %v", err)
		return AccountRegister, err
	} else {
		// Nếu insert thông tin user thành công vào DB rồi, thì cho vào queue gửi mail kích hoạt thanh cong ?
	}

	return &dataRes, nil
}
