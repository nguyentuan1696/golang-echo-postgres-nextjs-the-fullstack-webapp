package dto

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"time"
)

type TokenDetails struct {
	AccessToken        *jwt.Token
	RefreshToken       *jwt.Token
	SignedAccessToken  string
	SignedRefreshToken string
	Username           string
	RtExpires          time.Duration
	AtExpires          time.Duration
}

type AccountLoginRes struct {
	UserId       uuid.UUID  `json:"user_id"`
	Username     string     `json:"username"`
	AccessToken  *jwt.Token `json:"access_token"`
	RefreshToken *jwt.Token `json:"refresh_token"`
}

type AccountLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AccountRegister struct {
	Id        uuid.UUID `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	CreatedAt int64     `json:"created_at" db:"created_at"`
	UpdatedAt int64     `json:"updated_at" db:"created_at"`
}

type AccountRegisterRes struct {
	UserId   uuid.UUID `json:"user_id"`
	Username string    `json:"username"`
}
