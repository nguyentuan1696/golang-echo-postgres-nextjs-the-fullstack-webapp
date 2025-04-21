package router

import (
	"go-api-starter/core/middleware"
	"go-api-starter/modules/account/controller"

	"github.com/labstack/echo/v4"
)

type AccountRouter struct {
	controller *controller.AccountController
}

func NewAccountRouter(controller *controller.AccountController) *AccountRouter {
	return &AccountRouter{
		controller: controller,
	}
}

func (r *AccountRouter) Setup(e *echo.Echo, authMiddleware *middleware.Middleware) {
	// API v1 group
	v1 := e.Group("/api/v1")

	// Account routes group
	accounts := v1.Group("/accounts")

	// Auth routes - no middleware needed
	auth := accounts.Group("")
	auth.POST("/register", r.controller.Register)
	auth.POST("/login", r.controller.Login)
	auth.POST("/refresh-token", r.controller.RefreshToken)
	auth.POST("/forgot-password", r.controller.ForgotPassword)
	auth.POST("/reset-password", r.controller.ResetPassword)

	// User routes - requires authentication
	user := accounts.Group("")
	user.Use(authMiddleware.AuthMiddleware())
	user.POST("/logout", r.controller.Logout)
	user.PUT("/change-password", r.controller.ChangePassword)

	// Admin routes
	admin := v1.Group("/admin")
	admin.Use(authMiddleware.AuthMiddleware())

	// User management routes
	users := admin.Group("/users")
	users.Use(authMiddleware.PermissionMiddleware("read:users"))
	users.GET("", r.controller.GetUsers)

	// User admin routes
	userAdmin := users.Group("")
	userAdmin.Use(authMiddleware.PermissionMiddleware("write:users", "manage:users"))
	userAdmin.DELETE("/:id", r.controller.DeleteUser)
	userAdmin.POST("/:id/lock", r.controller.LockUser)
	userAdmin.POST("/:id/unlock", r.controller.UnlockUser)

	// RBAC management routes
	rbac := admin.Group("/rbac")
	rbac.Use(authMiddleware.PermissionMiddleware("manage:rbac"))
	rbac.GET("/roles", r.controller.GetRoles)
	rbac.POST("/roles", r.controller.CreateRole)
	rbac.GET("/permissions", r.controller.GetPermissions)
	rbac.POST("/permissions", r.controller.CreatePermission)
	rbac.POST("/roles/:roleId/permissions/:permissionId", r.controller.AssignPermissionToRole)
	rbac.POST("/roles/:roleId/users/:userId", r.controller.AssignRoleToUser)
}
