package models

import (
	"time"
)

// table "users"
type User struct {
	ID        int        `json:"id"`
	Username  string     `json:"username"`
	Password  string     `json:"-"`
	IsActive  bool       `json:"is_active"`
	VotedTo   string     `json:"voted_to" gorm:"-"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

type RegisterUserRequest struct {
	Username string `json:"username" validate:"required,max=50" example:"john (Max 50 Chars)"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Data           LoginResponseData `json:"data"`
	ResponseStatus ResponseStatus    `json:"status"`
}

type LoginResponseData struct {
	AccessToken string `json:"access_token"`
}

type ReportResponse struct {
	Data           ReportResponseData `json:"data"`
	ResponseStatus ResponseStatus     `json:"status"`
}

type ReportResponseData struct {
	VoteStatus VoteStatus                `json:"vote_status"`
	Users      []User                    `json:"users"`
	Candidates []ReportCandidateResponse `json:"candidates"`
}

type ReportCandidateResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      int    `json:"-"`
	VoteID      int    `json:"-"`
	VoteCount   int    `json:"vote_count" gorm:"-"`
}
