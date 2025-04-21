package entity

import (
	"thichlab-backend-slowpoke/core/entity"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID            uuid.UUID  `db:"id"`
	UserName      string     `db:"user_name"`
	Email         string     `db:"email"`
	Password      string     `db:"password"`
	IsSocialLogin bool       `db:"is_social_login"`
	IsLocked      bool       `db:"is_locked"`
	LockedAt      *time.Time `db:"locked_at"`
	LockReason    string     `db:"lock_reason"`
	UnlockedAt    *time.Time `db:"unlocked_at"`
	UnlockReason  string     `db:"unlock_reason"`
	CreatedAt     time.Time  `db:"created_at"`
	UpdatedAt     time.Time  `db:"updated_at"`
}

type PaginatedUsers = entity.Pagination[*User]

type UserProfile struct {
	UserId      uuid.UUID  `db:"user_id"`
	FullName    string     `db:"full_name"`
	Birthday    *time.Time `db:"birthday"`
	Gender      string     `db:"gender"`
	PhoneNumber string     `db:"phone_number"`
	Address     string     `db:"address"`
	AvatarUrl   string     `db:"avatar_url"`
	Bio         string     `db:"bio"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at"`
}

type Role struct {
	Id          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type Permission struct {
	Id          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type RolePermission struct {
	RoleId       uuid.UUID `db:"role_id"`
	PermissionId uuid.UUID `db:"permission_id"`
}

type UserRole struct {
	UserId uuid.UUID `db:"user_id"`
	RoleId uuid.UUID `db:"role_id"`
}
