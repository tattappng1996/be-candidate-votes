package models

import (
	"time"
)

type User struct {
	ID        int        `json:"id"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	IsActive  bool       `json:"is_active"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
