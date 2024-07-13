package dto

import "github.com/google/uuid"

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
