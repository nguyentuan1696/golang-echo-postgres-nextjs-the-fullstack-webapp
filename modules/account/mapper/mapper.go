package mapper

import (
	"go-api-starter/modules/account/dto"
	"go-api-starter/modules/account/entity"
)

func ToUserEntity(user *dto.CreateAccountRequest) *entity.User {

	if user == nil {
		return nil
	}

	return &entity.User{
		UserName: user.Username,
		Password: user.Password,
		Email:    user.Email,
	}
}

func ToPaginatedUsersResponse(users *entity.PaginatedUsers) *dto.PaginatedUsersResponse {
	if users == nil {
		return nil
	}

	userDtos := make([]*dto.UserResponse, 0, len(users.Items))
	for _, user := range users.Items {
		userDtos = append(userDtos, &dto.UserResponse{
			Id:            user.ID,
			Username:      user.UserName,
			Email:         user.Email,
			IsSocialLogin: user.IsSocialLogin,
			IsLocked:      user.IsLocked,
		})
	}

	return &dto.PaginatedUsersResponse{
		Items:       userDtos,
		TotalItems:  users.TotalItems,
		TotalPages:  users.TotalPages,
		CurrentPage: users.CurrentPage,
		PageSize:    users.PageSize,
	}
}

func ToRoleEntity(role *dto.CreateRoleRequest) *entity.Role {
	if role == nil {
		return nil
	}
	return &entity.Role{
		Name:        role.Name,
		Description: role.Description,
	}
}

func ToPermissionEntity(permission *dto.CreatePermissionRequest) *entity.Permission {
	if permission == nil {
		return nil
	}
	return &entity.Permission{
		Name:        permission.Name,
		Description: permission.Description,
	}
}

func ToRoleResponses(roles []*entity.Role) []*dto.RoleResponse {
	if roles == nil {
		return nil
	}
	roleResponses := make([]*dto.RoleResponse, 0, len(roles))
	for _, role := range roles {
		roleResponses = append(roleResponses, &dto.RoleResponse{
			Id:          role.Id,
			Name:        role.Name,
			Description: role.Description,
		})
	}
	return roleResponses
}

func ToPermissionResponses(permissions []*entity.Permission) []*dto.PermissionResponse {
	if permissions == nil {
		return nil
	}
	permissionResponses := make([]*dto.PermissionResponse, 0, len(permissions))
	for _, permission := range permissions {
		permissionResponses = append(permissionResponses, &dto.PermissionResponse{
			Id:          permission.Id,
			Name:        permission.Name,
			Description: permission.Description,
		})
	}
	return permissionResponses
}
