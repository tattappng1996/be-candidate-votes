package models

import "time"

// table votes
type Vote struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	CandidateID int       `json:"candidate_id"`
	CreatedAt   time.Time `json:"-"`
}

type VoteRequest struct {
	CandidateID int `json:"candidate_id" validate:"required"`
	UserID      int `json:"-"`
}
