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
	Name        string `json:"name" validate:"max=30" example:"karen (Max 30 Chars)"`
	Description string `json:"description"`
}

type CreateCandidateResponse struct {
	ResponseStatus ResponseStatus `json:"status"`
}

type UpdateCandidateRequest struct {
	ID          int    `json:"id" validate:"required"`
	Name        string `json:"name" validate:"max=30" example:"karen (Max 30 Chars)"`
	Description string `json:"description"`
}

type ListCandidateRequest struct {
	ID           int    `json:"-"` // for getbyID
	SearchByName string `json:"search_by_name"`
	Limit        int    `json:"limit"`
	Page         int    `json:"page"`
	PageCount    int    `json:"page_count"`
}

type ListCandidateResponse struct {
	Data           ListCandidateResponseData `json:"data"`
	ResponseStatus ResponseStatus            `json:"status"`
}

type ListCandidateResponseData struct {
	Candidates []CandidateResponse `json:"candidates"`
	Limit      int                 `json:"limit"`
	Page       int                 `json:"page"`
	PageCount  int                 `json:"page_count"`
}

type GetCandidateResponse struct {
	Data           CandidateResponse `json:"data"`
	ResponseStatus ResponseStatus    `json:"status"`
}

type CandidateResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	VoteCount   int    `json:"vote_count"`
}
