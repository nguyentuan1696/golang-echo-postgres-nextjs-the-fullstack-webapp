package validator

import (
	"go-api-starter/core/utils"
	"go-api-starter/core/validation"
	"go-api-starter/modules/account/dto"
)

func ValidateCreateAccount(dataRequest dto.CreateAccountRequest) *validation.ValidationResult {
	result := validation.NewValidationResult()

	// Validate email
	switch {
	case utils.IsEmpty(dataRequest.Email):
		result.AddError("email", "Email is required")
	case !utils.IsValidEmail(dataRequest.Email):
		result.AddError("email", "Invalid email format")
	case !utils.IsValidEmailDomain(dataRequest.Email):
		result.AddError("email", "Invalid email domain")
	}

	// Validate password and confirmation
	if len(dataRequest.Password) < 8 {
		result.AddError("password", "Password must be at least 8 characters")
	}
	if dataRequest.Password != dataRequest.ConfirmPassword {
		result.AddError("confirm_password", "Password and confirmation do not match")
	}

	// Validate username
	if utils.IsEmpty(dataRequest.Username) {
		result.AddError("username", "Username is required")
	}

	return result
}

func ValidateLogin(dataRequest dto.LoginRequest) *validation.ValidationResult {
	result := validation.NewValidationResult()

	// Validate email
	switch {
	case utils.IsEmpty(dataRequest.Email):
		result.AddError("email", "Email is required")
	case !utils.IsValidEmail(dataRequest.Email):
		result.AddError("email", "Invalid email format")
	case !utils.IsValidEmailDomain(dataRequest.Email):
		result.AddError("email", "Invalid email domain")
	}

	// Validate password
	if len(dataRequest.Password) < 8 {
		result.AddError("password", "Password must be at least 8 characters")
	}

	return result
}

func ValidateChangePassword(dataRequest dto.ChangePasswordRequest) *validation.ValidationResult {
	result := validation.NewValidationResult()
	// Validate old password
	if len(dataRequest.CurrentPassword) < 8 {
		result.AddError("old_password", "Old password must be at least 8 characters")
	}
	// Validate new password and confirmation
	if len(dataRequest.NewPassword) < 8 {
		result.AddError("new_password", "New password must be at least 8 characters")
	}
	if dataRequest.NewPassword != dataRequest.ConfirmPassword {
		result.AddError("confirm_password", "New password and confirmation do not match")
	}

	return result
}

func ValidateForgotPassword(dataRequest dto.ForgotPasswordRequest) *validation.ValidationResult {
	result := validation.NewValidationResult()
	// Validate email
	switch {
	case utils.IsEmpty(dataRequest.Email):
		result.AddError("email", "Email is required")
	case !utils.IsValidEmail(dataRequest.Email):
		result.AddError("email", "Invalid email format")
	case !utils.IsValidEmailDomain(dataRequest.Email):
		result.AddError("email", "Invalid email domain")
		return result
	}
	return result
}

func ValidateResetPassword(dataRequest dto.ResetPasswordRequest) *validation.ValidationResult {
	result := validation.NewValidationResult()
	// Validate new password and confirmation
	if len(dataRequest.NewPassword) < 8 {
		result.AddError("new_password", "New password must be at least 8 characters")
	}
	if dataRequest.NewPassword != dataRequest.ConfirmPassword {
		result.AddError("confirm_password", "New password and confirmation do not match")
	}
	return result
}

func ValidateCreateRole(dataRequest *dto.CreateRoleRequest) *validation.ValidationResult {
	if dataRequest == nil {
		return nil
	}

	result := validation.NewValidationResult()
	// Validate name
	if utils.IsEmpty(dataRequest.Name) {
		result.AddError("name", "Name is required")
	}
	return result
}

func ValidateCreatePermission(dataRequest *dto.CreatePermissionRequest) *validation.ValidationResult {
	if dataRequest == nil {
		return nil
	}
	result := validation.NewValidationResult()
	// Validate name
	if utils.IsEmpty(dataRequest.Name) {
		result.AddError("name", "Name is required")
	}
	return result
}

func ValidateAssignPermissionToRole(dataRequest *dto.AssignPermissionToRoleRequest) *validation.ValidationResult {
	if dataRequest == nil {
		return nil
	}
	result := validation.NewValidationResult()
	// Validate role_id
	if utils.IsEmpty(utils.ToString(dataRequest.RoleId)) {
		result.AddError("role_id", "Role ID is required")
	}
	// Validate permission_id
	if utils.IsEmpty(utils.ToString(dataRequest.PermissionId)) {
		result.AddError("permission_id", "Permission ID is required")
	}
	return result
}

func ValidateAssignRoleToUser(dataRequest *dto.AssignRoleToUserRequest) *validation.ValidationResult {
	if dataRequest == nil {
		return nil
	}
	result := validation.NewValidationResult()
	// Validate user_id
	if utils.IsEmpty(utils.ToString(dataRequest.UserId)) {
		result.AddError("user_id", "User ID is required")
	}
	// Validate role_id
	if utils.IsEmpty(utils.ToString(dataRequest.RoleId)) {
		result.AddError("role_id", "Role ID is required")
	}
	return result
}
