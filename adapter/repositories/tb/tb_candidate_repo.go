package tb

import (
	"be-cadidate-votes/models"
	"context"
	"fmt"

	"gorm.io/gorm"
)

func (r *tbRepo) ListCandidateWithVote(ctx context.Context, filter models.ListCandidateRequest) ([]models.CandidateResponse, error) {
	query := r.db.WithContext(ctx).
		Select(`c.*, COUNT(v.id) AS vote_count`).Table(`candidates c`).
		Joins(`LEFT JOIN votes v ON v.candidate_id = c.id`).
		Where(`c.deleted_at IS NULL`).Group(`c.id`)

	if filter.ID > 0 {
		query = query.Where(`c.id = ?`, filter.ID)
	}
	if filter.SearchByName != "" {
		query = query.Where(`c.name ILIKE ?`, fmt.Sprintf("%%%s%%", filter.SearchByName))
	}
	if filter.OrderByVoteCount {
		query = query.Order(`vote_count DESC`)
	}
	if filter.Limit > 0 {
		query = query.Limit(filter.Limit)
	}
	if filter.Offset > 0 {
		query = query.Offset(filter.Offset)
	}

	c := []models.CandidateResponse{}
	if err := query.Scan(&c).Error; err != nil {
		return c, err
	}

	return c, nil
}

func (r *tbRepo) CountCandidate(ctx context.Context, filter models.ListCandidateRequest) (int, error) {
	query := r.db.WithContext(ctx).Table(`candidates c`).Where(`c.deleted_at IS NULL`)

	if filter.SearchByName != "" {
		query = query.Where(`c.name ILIKE ?`, fmt.Sprintf("%%%s%%", filter.SearchByName))
	}

	var count int64
	if err := query.Count(&count).Error; err != nil {
		return int(count), err
	}

	return int(count), nil
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
