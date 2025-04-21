package service

import (
	"context"
	"fmt"
	"go-api-starter/core/constants"
	"go-api-starter/core/errors"
	"go-api-starter/core/logger"
	"go-api-starter/core/utils"
	"go-api-starter/modules/account/dto"
	"go-api-starter/modules/account/mapper"
	"time"

	"github.com/google/uuid"
)

func (s *AccountService) CreateAccount(ctx context.Context, requestData *dto.CreateAccountRequest) (*dto.CreateAccountResponse, *errors.AppError) {

	existingUser, err := s.repo.GetUserByEmailOrUserNameOrId(ctx, requestData.Email, requestData.Username, uuid.Nil)
	if err != nil {
		logger.Error("AccountService:CreateAccount:Failed to check existing user", "error", err)
		return nil, errors.NewAppError(errors.ErrAlreadyExists, "", err)
	}
	if existingUser != nil {
		return nil, errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(requestData.Password)
	if err != nil {
		logger.Error("AccountService:CreateAccount:Failed to hash password", "error", err)
		return nil, errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}

	// Convert DTO to entity
	user := mapper.ToUserEntity(requestData)
	user.Password = hashedPassword

	// Save to database
	createdUser, err := s.repo.CreateAccount(ctx, user)
	if err != nil {
		logger.Error("AccountService:CreateAccount:Failed to create account", "error", err, "email", requestData.Email)
		return nil, errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}

	// Generate access token (expires in 1 day)
	accessToken, err := utils.GenerateToken(createdUser.ID, createdUser.Email, createdUser.UserName, constants.AccessTokenExpiry)
	if err != nil {
		logger.Error("AccountService:CreateAccount:Failed to generate access token", "error", err)
		return nil, errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}

	// Generate refresh token (expires in 7 days)
	refreshToken, err := utils.GenerateToken(createdUser.ID, createdUser.Email, createdUser.UserName, constants.RefreshTokenExpiry)
	if err != nil {
		logger.Error("AccountService:CreateAccount:Failed to generate refresh token", "error", err)
		return nil, errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}

	// Prepare response
	response := &dto.CreateAccountResponse{
		Username:     createdUser.UserName,
		Email:        createdUser.Email,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return response, nil
}

// TODO: Improve login logic
func (s *AccountService) Login(ctx context.Context, requestData *dto.LoginRequest) (*dto.LoginResponse, *errors.AppError) {
	// Check rate limiting first
	rateKey := fmt.Sprintf("login_rate:%s", requestData.Email)
	rateCount, _ := s.cache.Get(ctx, rateKey).Int()
	if rateCount >= 10 { // Max 10 attempts per minute
		return nil, errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", nil)
	}

	// Increment and set rate limit
	s.cache.Incr(ctx, rateKey)
	s.cache.Expire(ctx, rateKey, time.Minute) // Reset after 1 minute

	// Check if user is blocked
	blockKey := fmt.Sprintf("login_blocked:%s", requestData.Email)
	blocked, err := s.cache.IsLoginBlocked(ctx, blockKey)
	if err != nil {
		return nil, errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}
	if blocked {
		return nil, errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}

	existingUser, err := s.repo.GetUserByEmailOrUserNameOrId(ctx, requestData.Email, "", uuid.Nil)
	if err != nil {
		logger.Error("AccountService:Login:Failed to check existing user", "error", err)
		return nil, errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}

	// Check credentials
	if existingUser == nil {
		s.cache.Incr(ctx, blockKey)
		s.cache.Expire(ctx, blockKey, time.Hour)
		return nil, errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}

	if !utils.ComparePassword(existingUser.Password, requestData.Password) {
		s.cache.Incr(ctx, blockKey)
		s.cache.Expire(ctx, blockKey, time.Hour)
		return nil, errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}

	// Success - clear all rate limiting and blocking
	s.cache.Del(ctx, rateKey)
	s.cache.Del(ctx, blockKey)

	// Generate access token (expires in 1 day)
	accessToken, err := utils.GenerateToken(existingUser.ID, existingUser.Email, existingUser.UserName, 24*time.Duration(time.Hour))
	if err != nil {
		logger.Error("AccountService:Login:Failed to generate access token", "error", err)
		return nil, errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}
	// Generate refresh token (expires in 7 days)
	refreshToken, err := utils.GenerateToken(existingUser.ID, existingUser.Email, existingUser.UserName, 7*24*time.Duration(time.Hour))
	if err != nil {
		logger.Error("AccountService:Login:Failed to generate refresh token", "error", err)
		return nil, errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}
	// Prepare response
	response := &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return response, nil
}

func (s *AccountService) ChangePassword(ctx context.Context, token string, requestData *dto.ChangePasswordRequest) *errors.AppError {
	// Validate and parse token
	claims, err := utils.ValidateAndParseToken(token)
	if err != nil {
		logger.Error("AccountService:ChangePassword:Failed to validate token", "error", err)
		return errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}

	// Get user from database
	user, err := s.repo.GetUserByEmailOrUserNameOrId(ctx, "", "", claims.UserID)
	if err != nil {
		logger.Error("AccountService:ChangePassword:Failed to get user", "error", err)
		return errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}
	if user == nil {
		return errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}

	// Verify current password
	if !utils.ComparePassword(user.Password, requestData.CurrentPassword) {
		return errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}

	// Hash the new password
	hashedPassword, err := utils.HashPassword(requestData.NewPassword)
	if err != nil {
		logger.Error("AccountService:ChangePassword:Failed to hash new password", "error", err)
		return errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}

	// Update password in database
	user.Password = hashedPassword
	err = s.repo.UpdatePassword(ctx, user)
	if err != nil {
		logger.Error("AccountService:ChangePassword:Failed to update password", "error", err)
		return errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}

	return nil
}

func (s *AccountService) ForgotPassword(ctx context.Context, requestData *dto.ForgotPasswordRequest) *errors.AppError {
	// Rate limiting check
	attemptsKey := fmt.Sprintf("forgot_password_attempts:%s", requestData.Email)
	attempts, err := s.cache.Get(ctx, attemptsKey).Int()
	if err == nil && attempts >= 3 {
		return errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}

	// Increment attempt counter
	s.cache.Incr(ctx, attemptsKey)
	s.cache.Expire(ctx, attemptsKey, time.Hour)

	// Get user from database
	user, err := s.repo.GetUserByEmailOrUserNameOrId(ctx, requestData.Email, "", uuid.Nil)
	if err != nil {
		logger.Error("AccountService:ForgotPassword:Failed to get user", "error", err)
		return errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}

	// Check if user exists
	if user == nil {
		return errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}

	// Generate reset token (expires in 15 minutes)
	resetToken, err := utils.GenerateToken(user.ID, user.Email, user.UserName, 15*time.Minute)
	if err != nil {
		logger.Error("AccountService:ForgotPassword:Failed to generate reset token", "error", err)
		return errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}

	fmt.Printf("Reset token: %s\n", resetToken)

	// Create reset password link
	// fmt.Sprintf("%s/reset-password?token=%s", "http://localhost:3000", resetToken)

	// Prepare email content

	// Send email

	return errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
}

func (s *AccountService) ResetPassword(ctx context.Context, requestData *dto.ResetPasswordRequest) *errors.AppError {
	claims, err := utils.ValidateAndParseToken(requestData.ResetToken)
	if err != nil {
		logger.Error("AccountService:ResetPassword:Invalid token", "error", err)
		return errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}

	user, err := s.repo.GetUserByEmailOrUserNameOrId(ctx, "", "", claims.UserID)
	if err != nil {
		logger.Error("AccountService:ResetPassword:Failed to get user", "error", err)
		return errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}
	if user == nil {
		return errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}

	hashedPassword, err := utils.HashPassword(requestData.NewPassword)
	if err != nil {
		logger.Error("AccountService:ResetPassword:Failed to hash password", "error", err)
		return errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}

	user.Password = hashedPassword
	if err = s.repo.UpdatePassword(ctx, user); err != nil {
		logger.Error("AccountService:ResetPassword:Failed to update password", "error", err)
		return errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}

	return nil
}

func (s *AccountService) Logout(ctx context.Context, token string) *errors.AppError {
	// Add token to blacklist with expiry matching the token's expiry
	claims, err := utils.ValidateAndParseToken(token)
	if err != nil {
		logger.Error("AccountService:Logout:Failed to validate token", "error", err)
		return errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}

	// Calculate remaining time until token expiry
	expiryTime := time.Until(time.Unix(claims.ExpiresAt.Unix(), 0)).Seconds()
	if expiryTime <= 0 {
		return nil // Token already expired
	}

	// Add to blacklist
	blacklistKey := fmt.Sprintf("token_blacklist:%s", token)
	err = s.cache.Set(ctx, blacklistKey, true, time.Duration(expiryTime)*time.Second)
	if err != nil {
		logger.Error("AccountService:Logout:Failed to blacklist token", "error", err)
		return errors.NewAppError(errors.ErrAlreadyExists, "AccountService:CreateAccount:username or email already exists", err)
	}

	return nil
}
