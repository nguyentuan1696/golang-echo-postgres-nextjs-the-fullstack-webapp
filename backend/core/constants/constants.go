package constants

import "time"

const (
	AccessTokenExpiry  = 24 * time.Hour
	RefreshTokenExpiry = 7 * 24 * time.Hour
)

// Giới hạn login
const (
	MaxLoginAttempts = 5
	BlockDuration    = 15 * time.Minute
)
