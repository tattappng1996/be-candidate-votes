package tb

import (
	"be-cadidate-votes/models"
	"context"

	"gorm.io/gorm"
)

type TbRepo interface {
	GetUser(ctx context.Context, filter models.User) (models.User, error)
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)

	ListCandidateWithVote(ctx context.Context, filter models.ListCandidateRequest) ([]models.CandidateResponse, error)
	CountCandidate(ctx context.Context, filter models.ListCandidateRequest) (int, error)
	CreateCandidate(ctx context.Context, user *models.Candidate) (*models.Candidate, error)
	UpdateCandidate(ctx context.Context, user models.Candidate) error
}

func NewTbRepository(db *gorm.DB) TbRepo {
	return &tbRepo{
		db: db,
	}
}

type tbRepo struct {
	db *gorm.DB
}
