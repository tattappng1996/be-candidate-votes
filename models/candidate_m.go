package models

import "time"

const (
	DefaultLimit = 50
	DefaultPage  = 1
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
	ID               int    `json:"-"` // for getbyID
	SearchByName     string `json:"search_by_name"`
	OrderByVoteCount bool   `json:"order_by_vote_count"`
	Limit            int    `json:"limit"`
	Page             int    `json:"page"`
	Offset           int    `json:"-"`
}

type ListCandidateResponse struct {
	Data           ListCandidateResponseData `json:"data"`
	ResponseStatus ResponseStatus            `json:"status"`
}

type ListCandidateResponseData struct {
	Candidates     []CandidateResponse `json:"candidates"`
	CandidateCount int                 `json:"candidate_count"`
	Limit          int                 `json:"limit"`
	Page           int                 `json:"page"`
	PageCount      int                 `json:"page_count"`
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
