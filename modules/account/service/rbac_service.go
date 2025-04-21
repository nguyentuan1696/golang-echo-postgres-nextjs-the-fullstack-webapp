package service

import (
	"context"

	"go-api-starter/core/errors"
	"go-api-starter/modules/account/dto"
	"go-api-starter/modules/account/mapper"
	"time"

	"github.com/google/uuid"
)

func (s *AccountService) CreateRole(ctx context.Context, role *dto.CreateRoleRequest) *errors.AppError {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := s.repo.CreateRole(ctx, mapper.ToRoleEntity(role))
	if err != nil {
		return errors.NewAppError(errors.ErrInternal, "AccountService:CreateAccount:username or email already exists", err)
	}
	return nil
}

func (s *AccountService) CreatePermission(ctx context.Context, permission *dto.CreatePermissionRequest) *errors.AppError {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	err := s.repo.CreatePermission(ctx, mapper.ToPermissionEntity(permission))
	if err != nil {
		return errors.NewAppError(errors.ErrInternal, "AccountService:CreateAccount:username or email already exists", err)
	}
	return nil
}

func (s *AccountService) GetPermissions(ctx context.Context) ([]*dto.PermissionResponse, *errors.AppError) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	permissions, err := s.repo.GetPermissions(ctx)
	if err != nil {
		return nil, errors.NewAppError(errors.ErrInternal, "AccountService:GetPermissions:internal server error", err)
	}
	return mapper.ToPermissionResponses(permissions), nil
}

func (s *AccountService) GetRoles(ctx context.Context) ([]*dto.RoleResponse, *errors.AppError) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	roles, err := s.repo.GetRoles(ctx)
	if err != nil {
		return nil, errors.NewAppError(errors.ErrInternal, "AccountService:GetRoles:internal server error", err)
	}
	return mapper.ToRoleResponses(roles), nil
}

func (s *AccountService) AssignRoleToUser(ctx context.Context, userID uuid.UUID, roleID uuid.UUID) *errors.AppError {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	userExists, err := s.repo.GetUserByEmailOrUserNameOrId(ctx, "", "", userID)
	if err != nil {
		return errors.NewAppError(errors.ErrInternal, "AccountService:AssignRoleToUser:internal server error", err)
	}
	if userExists == nil {
		return errors.NewAppError(errors.ErrNotFound, "AccountService:AssignRoleToUser:user not found", nil)
	}

	roleExists, err := s.repo.RoleExists(ctx, roleID)
	if err != nil {
		return errors.NewAppError(errors.ErrInternal, "AccountService:AssignRoleToUser:internal server error", err)
	}
	if !roleExists {
		return errors.NewAppError(errors.ErrNotFound, "AccountService:AssignRoleToUser:role not found", nil)
	}

	errAssignRoleToUser := s.repo.AssignRoleToUser(ctx, userID, roleID)
	if errAssignRoleToUser != nil {
		return errors.NewAppError(errors.ErrInternal, "AccountService:AssignRoleToUser:internal server error", errAssignRoleToUser)
	}
	return nil
}

func (s *AccountService) AssignPermissionToRole(ctx context.Context, roleID uuid.UUID, permissionID uuid.UUID) *errors.AppError {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	userExists, err := s.repo.GetUserByEmailOrUserNameOrId(ctx, "", "", uuid.Nil)
	if err != nil {
		return errors.NewAppError(errors.ErrInternal, "AccountService:AssignPermissionToRole:internal server error", err)
	}
	if userExists == nil {
		return errors.NewAppError(errors.ErrNotFound, "AccountService:AssignPermissionToRole:user not found", nil)
	}

	roleExists, err := s.repo.RoleExists(ctx, roleID)
	if err != nil {
		return errors.NewAppError(errors.ErrInternal, "AccountService:AssignPermissionToRole:internal server error", err)
	}
	if !roleExists {
		return errors.NewAppError(errors.ErrNotFound, "AccountService:AssignPermissionToRole:role not found", nil)
	}
	permissionExists, err := s.repo.PermissionExists(ctx, permissionID)
	if err != nil {
		return errors.NewAppError(errors.ErrInternal, "AccountService:AssignPermissionToRole:internal server error", err)
	}
	if !permissionExists {
		return errors.NewAppError(errors.ErrNotFound, "AccountService:AssignPermissionToRole:permission not found", nil)
	}

	errAssignPermissionToRole := s.repo.AssignPermissionToRole(ctx, roleID, permissionID)
	if errAssignPermissionToRole != nil {
		return errors.NewAppError(errors.ErrInternal, "AccountService:AssignPermissionToRole:internal server error", errAssignPermissionToRole)
	}
	return nil
}

func (s *AccountService) HasPermission(ctx context.Context, userID uuid.UUID, permissionID uuid.UUID) (bool, *errors.AppError) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	userExists, err := s.repo.GetUserByEmailOrUserNameOrId(ctx, "", "", userID)
	if err != nil {
		return false, errors.NewAppError(errors.ErrInternal, "AccountService:HasPermission:internal server error", err)
	}
	if userExists == nil {
		return false, errors.NewAppError(errors.ErrNotFound, "AccountService:HasPermission:user not found", nil)
	}
	permissionExists, err := s.repo.PermissionExists(ctx, permissionID)
	if err != nil {
		return false, errors.NewAppError(errors.ErrInternal, "AccountService:HasPermission:internal server error", err)
	}
	if !permissionExists {
		return false, errors.NewAppError(errors.ErrNotFound, "AccountService:HasPermission:permission not found", nil)
	}
	hasPermission, err := s.repo.HasPermission(ctx, userID, permissionID)
	if err != nil {
		return false, errors.NewAppError(errors.ErrInternal, "AccountService:HasPermission:internal server error", err)
	}
	return hasPermission, nil
}
