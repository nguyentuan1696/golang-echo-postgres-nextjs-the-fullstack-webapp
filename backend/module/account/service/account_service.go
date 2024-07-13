package service

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"strings"
	"thichlab-backend-docs/dto"
	"thichlab-backend-docs/infrastructure/logger"
	"thichlab-backend-docs/infrastructure/util"
	"time"
)

func (service *AccountService) AccountLogin(ctx context.Context, args *dto.AccountLoginReq) (*dto.AccountLoginRes, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Compare password
	infoUser, _ := service.AccountRepository.GetInfoUserByUsername(ctx, args.Username)
	if !util.CheckPasswordHash(args.Password, infoUser.Password) {
		return nil, nil
	}

	// Tao Session, Tao Access Token va Refresh Token
	td := new(dto.TokenDetails)
	td.Username = args.Username

	accessTokenClaims := jwt.MapClaims{}
	accessTokenClaims["username"] = args.Username
	accessTokenClaims["user_id"] = infoUser.Id
	td.AccessToken = jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)

	refreshTokenClaims := jwt.MapClaims{}
	refreshTokenClaims["username"] = args.Username
	refreshTokenClaims["user_id"] = infoUser.Id
	td.RefreshToken = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	signKey := viper.GetString("OpenIDJwt.SecretKey")
	_, _, err := util.Sign(td, signKey)

	return nil, err

}

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
