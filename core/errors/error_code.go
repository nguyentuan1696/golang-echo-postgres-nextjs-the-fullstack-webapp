package errors

type ErrorCode int

const (
	// Authentication & Authorization errors (1000-1999)
	ErrInvalidCredentials ErrorCode = 1000
	ErrTokenExpired       ErrorCode = 1001
	ErrUnauthorized       ErrorCode = 1002
	ErrForbidden          ErrorCode = 1003

	// Input validation errors (2000-2999)
	ErrInvalidInput    ErrorCode = 2000
	ErrInvalidEmail    ErrorCode = 2001
	ErrInvalidPassword ErrorCode = 2002
	ErrInvalidFormat   ErrorCode = 2003

	// Resource errors (3000-3999)
	ErrNotFound        ErrorCode = 3000
	ErrAlreadyExists   ErrorCode = 3001
	ErrResourceLocked  ErrorCode = 3002
	ErrResourceExpired ErrorCode = 3003

	// Database errors (4000-4999)
	ErrDatabase        ErrorCode = 4000
	ErrDatabaseTimeout ErrorCode = 4001
	ErrUniqueViolation ErrorCode = 4002
	ErrForeignKey      ErrorCode = 4003

	// Business logic errors (5000-5999)
	ErrBusinessRule    ErrorCode = 5000
	ErrInvalidState    ErrorCode = 5001
	ErrLimitExceeded   ErrorCode = 5002
	ErrOperationFailed ErrorCode = 5003

	// System errors (6000-6999)
	ErrInternal      ErrorCode = 6000
	ErrConfiguration ErrorCode = 6001
	ErrThirdParty    ErrorCode = 6002
	ErrNetwork       ErrorCode = 6003
)


