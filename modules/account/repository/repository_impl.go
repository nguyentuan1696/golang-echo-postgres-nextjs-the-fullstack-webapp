package repository

import (
	"context"
	"go-api-starter/core/database"
	"go-api-starter/modules/account/entity"

	"github.com/google/uuid"
)

type AccountRepository struct {
	DB database.Database
}

func NewAccountRepository(db database.Database) IAccountRepository {
	return &AccountRepository{
		DB: db,
	}
}

type IAccountRepository interface {
	GetUserByEmailOrUserNameOrId(ctx context.Context, email, userName string, userId uuid.UUID) (*entity.User, error)
	CreateAccount(ctx context.Context, user *entity.User) (*entity.User, error)
	UpdatePassword(ctx context.Context, user *entity.User) error
	GetUsers(ctx context.Context, pageNumber, pageSize int) (*entity.PaginatedUsers, error)

	// Rbac
	CreateRole(ctx context.Context, role *entity.Role) error
	GetRoles(ctx context.Context) ([]*entity.Role, error)
	CreatePermission(ctx context.Context, permission *entity.Permission) error
	GetPermissions(ctx context.Context) ([]*entity.Permission, error)
	AssignPermissionToRole(ctx context.Context, roleID uuid.UUID, permissionID uuid.UUID) error
	AssignRoleToUser(ctx context.Context, userID uuid.UUID, roleID uuid.UUID) error
	RoleExists(ctx context.Context, roleID uuid.UUID) (bool, error)
	PermissionExists(ctx context.Context, permissionID uuid.UUID) (bool, error)
	HasPermission(ctx context.Context, userID uuid.UUID, permissionID uuid.UUID) (bool, error)
	DeleteRole(ctx context.Context, roleID uuid.UUID) error
	DeletePermission(ctx context.Context, permissionID uuid.UUID) error
}
