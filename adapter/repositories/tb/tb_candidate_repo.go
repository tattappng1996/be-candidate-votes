package tb

import (
	"be-cadidate-votes/models"
	"context"

	"gorm.io/gorm"
)

func (r *tbRepo) ListCandidateWithVote(ctx context.Context, filter models.ListCandidateRequest) ([]models.CandidateResponse, error) {
	query := r.db.WithContext(ctx).
		Select(`c.*, COUNT(v.id) AS vote_count`).Table(`candidates c`).
		Joins(`LEFT JOIN votes v ON v.candidate_id = c.id`).Group(`c.id`)

	if filter.ID > 0 {
		query = query.Where(`c.id = ?`, filter.ID)
	}

	c := []models.CandidateResponse{}
	if err := query.Scan(&c).Error; err != nil {
		return c, err
	}

	return c, nil
}

func (r *tbRepo) CreateCandidate(ctx context.Context, candidate *models.Candidate) (*models.Candidate, error) {
	result := r.db.WithContext(ctx).Table(`candidates`).Create(candidate)
	if result.Error != nil {
		return nil, result.Error
	}
	return candidate, nil
}

func (r *tbRepo) UpdateCandidate(ctx context.Context, candidate models.Candidate) error {
	result := r.db.WithContext(ctx).Table(`candidates`).Updates(candidate)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
