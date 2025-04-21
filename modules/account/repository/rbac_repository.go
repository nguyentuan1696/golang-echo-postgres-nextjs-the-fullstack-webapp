package repository

import (
	"context"
	"go-api-starter/core/logger"
	"go-api-starter/modules/account/entity"

	"github.com/google/uuid"
)

func (r *AccountRepository) CreateRole(ctx context.Context, role *entity.Role) error {
	_, err := r.DB.NamedExecContext(ctx, `INSERT INTO roles (name, description) VALUES (:name, :description)`, role)
	logger.Error("AccountRepository:CreateRole:", err)
	return err
}

func (r *AccountRepository) GetRoles(ctx context.Context) ([]*entity.Role, error) {
	var roles []*entity.Role
	err := r.DB.SelectContext(ctx, &roles, `SELECT * FROM roles`)
	if err != nil {
		logger.Error("AccountRepository:GetRoles:", err)
		return nil, err
	}
	return roles, err
}

func (r *AccountRepository) CreatePermission(ctx context.Context, permission *entity.Permission) error {
	_, err := r.DB.NamedExecContext(ctx, `INSERT INTO permissions (name, description) VALUES (:name, :description)`, permission)
	if err != nil {
		logger.Error("AccountRepository:CreatePermission:", err)
		return err
	}
	return err
}

func (r *AccountRepository) GetPermissions(ctx context.Context) ([]*entity.Permission, error) {
	var permissions []*entity.Permission
	err := r.DB.SelectContext(ctx, &permissions, `SELECT * FROM permissions`)
	if err != nil {
		logger.Error("AccountRepository:GetPermissions:", err)
		return nil, err
	}
	return permissions, err
}

func (r *AccountRepository) AssignPermissionToRole(ctx context.Context, roleID uuid.UUID, permissionID uuid.UUID) error {
	err := r.DB.ExecContext(ctx, `INSERT INTO role_permissions (role_id, permission_id) VALUES ($1, $2)`, roleID, permissionID)
	logger.Error("AccountRepository:AssignPermissionToRole:", err)
	return err
}

func (r *AccountRepository) AssignRoleToUser(ctx context.Context, userID uuid.UUID, roleID uuid.UUID) error {
	err := r.DB.ExecContext(ctx, `INSERT INTO user_roles (user_id, role_id) VALUES ($1, $2)`, userID, roleID)
	logger.Error("AccountRepository:AssignRoleToUser:", err)
	return err
}

func (r *AccountRepository) RoleExists(ctx context.Context, roleID uuid.UUID) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM roles WHERE id = $1)`

	err := r.DB.GetContext(ctx, &exists, query, roleID)
	if err != nil {
		logger.Error("AccountRepository:RoleExists:", err)
		return false, err
	}

	return exists, nil
}

func (r *AccountRepository) PermissionExists(ctx context.Context, permissionID uuid.UUID) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM permissions WHERE id = $1)`

	err := r.DB.GetContext(ctx, &exists, query, permissionID)
	if err != nil {
		logger.Error("AccountRepository:PermissionExists:", err)
		return false, err
	}

	return exists, nil
}

func (r *AccountRepository) DeleteRole(ctx context.Context, roleID uuid.UUID) error {
	err := r.DB.ExecContext(ctx, `DELETE FROM roles WHERE id = $1`, roleID)
	if err != nil {
		logger.Error("AccountRepository:DeleteRole:", err)
	}
	return err
}

func (r *AccountRepository) DeletePermission(ctx context.Context, permissionID uuid.UUID) error {
	err := r.DB.ExecContext(ctx, `DELETE FROM permissions WHERE id = $1`, permissionID)
	if err != nil {
		logger.Error("AccountRepository:DeletePermission:", err)
	}

	return err
}

func (r *AccountRepository) HasPermission(ctx context.Context, userID uuid.UUID, permissionID uuid.UUID) (bool, error) {
	var exists bool
	query := `
		SELECT 1 FROM user_roles ur
	JOIN role_permissions rp ON ur.role_id = rp.role_id
	JOIN permissions p ON rp.permission_id = p.id
	WHERE ur.user_id = $1 AND p.name = $2
	LIMIT 1
	`
	err := r.DB.GetContext(ctx, &exists, query, userID, permissionID)
	if err != nil {
		logger.Error("AccountRepository:HasPermission:", err)
		return false, err
	}
	return exists, nil
}
