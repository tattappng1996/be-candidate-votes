package models

import "time"

// table votes
type Vote struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	CandidateID int       `json:"candidate_id"`
	CreatedAt   time.Time `json:"-"`
}
