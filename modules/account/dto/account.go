package dto

import (
	"thichlab-backend-slowpoke/core/dto"

	"github.com/google/uuid"
)

type CreateAccountRequest struct {
	Username        string `json:"user_name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type CreateAccountResponse struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

type ResetPasswordRequest struct {
	ResetToken      string `json:"reset_token"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

type UserResponse struct {
	Id            uuid.UUID `json:"id"`
	Username      string    `json:"username"`
	Email         string    `json:"email"`
	IsSocialLogin bool      `json:"is_social_login"`
	IsLocked      bool      `json:"is_locked"`
}

type PaginatedUsersResponse = dto.Pagination[*UserResponse]

type CreateRoleRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type RoleResponse struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type CreatePermissionRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type PermissionResponse struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type AssignPermissionToRoleRequest struct {
	RoleId       uuid.UUID `json:"role_id"`
	PermissionId uuid.UUID `json:"permission_id"`
}
type AssignRoleToUserRequest struct {
	UserId uuid.UUID `json:"user_id"`
	RoleId uuid.UUID `json:"role_id"`
}
