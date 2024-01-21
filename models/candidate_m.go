package models

import "time"

const (
	DefaultCandidateDescription = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
)

// table candidates
type Candidate struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-"`
}

type CreateCandidateRequest struct {
	Name        string `json:"name" validate:"required,max=30" example:"karen (Max 30 Chars)"`
	Description string `json:"description"`
}

type CreateCandidateResponse struct {
	ResponseStatus ResponseStatus `json:"status"`
}
