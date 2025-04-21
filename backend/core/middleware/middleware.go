package middleware

import (
	"net/http"
	"strings"
	"thichlab-backend-slowpoke/core/controller"
	"thichlab-backend-slowpoke/core/logger"
	"thichlab-backend-slowpoke/core/utils"
	"thichlab-backend-slowpoke/modules/account/service"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Middleware struct {
	controller.BaseController
	accountService service.IAccountService
}

func NewMiddleware(accountService service.IAccountService) *Middleware {
	return &Middleware{
		BaseController: controller.NewBaseController(),
		accountService: accountService,
	}
}
func (m *Middleware) AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get token from header
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return m.Unauthorized("missing authorization header")
			}

			// Check Bearer token format
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return m.Unauthorized("missing authorization header")
			}

			// Validate token
			claims, err := utils.ValidateToken(parts[1])
			if err != nil {
				return m.Unauthorized("invalid token")
			}

			// Set user claims in context
			c.Set("user", claims)
			return next(c)
		}
	}
}

func (m *Middleware) PermissionMiddleware(requiredPermissions ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userClaims, ok := c.Get("user").(*utils.Claims)
			if !ok {
				return m.Unauthorized("missing authorization header")
			}

			// Use account service to check permissions
			hasPermission, err := m.accountService.HasPermission(c.Request().Context(), userClaims.UserID, uuid.Nil)
			if err != nil {
				logger.Error("Error checking permissions", "error", err)
				return m.InternalServerError("error checking permissions")
			}

			if !hasPermission {
				return m.Forbidden("insufficient permissions")
			}

			return next(c)
		}
	}
}

func CORSMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Access-Control-Allow-Origin", "*")
			c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Response().Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
			c.Response().Header().Set("Access-Control-Allow-Credentials", "true")

			if c.Request().Method == "OPTIONS" {
				return c.NoContent(http.StatusOK)
			}

			return next(c)
		}
	}
}

func LoggerMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()

			if err := next(c); err != nil {
				c.Error(err)
			}

			// Log request details
			logger.Info("Request",
				"method", req.Method,
				"uri", req.RequestURI,
				"status", res.Status,
				"remote_ip", c.RealIP(),
			)

			return nil
		}
	}
}
