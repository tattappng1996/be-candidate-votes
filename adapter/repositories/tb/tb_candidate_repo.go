package tb

import (
	"be-cadidate-votes/models"
	"context"
)

func (r *tbRepo) GetCandidate(ctx context.Context, filter models.Candidate) (models.Candidate, error) {
	query := r.db.WithContext(ctx).Table(`candidates`)

	if filter.ID > 0 {
		query = query.Where(`id = ?`, filter.ID)
	}
	if filter.Name != "" {
		query = query.Where(`name = ?`, filter.Name)
	}

	c := models.Candidate{}
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
