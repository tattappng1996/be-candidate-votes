package models

import (
	"time"
)

// table "users"
type User struct {
	ID        int        `json:"id"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	IsActive  bool       `json:"is_active"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

type RegisterUserRequest struct {
	Username string `json:"username" validate:"required,max=50" example:"john (Max 50 Chars)"`
	Password string `json:"password" validate:"required"`
}

type RegisterUserResponse struct {
	ResponseStatus ResponseStatus `json:"status"`
}

type LoginResponse struct {
	Data           LoginResponseData `json:"data"`
	ResponseStatus ResponseStatus    `json:"status"`
}

type LoginResponseData struct {
	AccessToken string `json:"access_token"`
}
