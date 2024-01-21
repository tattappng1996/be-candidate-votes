package tb

import (
	"be-cadidate-votes/models"
	"context"

	"gorm.io/gorm"
)

func (r *tbRepo) GetVoteStatus(ctx context.Context) (models.VoteStatus, error) {
	query := r.db.WithContext(ctx).Table(`vote_statuses`)

	v := models.VoteStatus{}
	if err := query.Scan(&v).Error; err != nil {
		return v, err
	}

	return v, nil
}

func (r *tbRepo) UpdateVoteStatus(ctx context.Context, vs models.VoteStatus) error {
	result := r.db.WithContext(ctx).Exec("UPDATE vote_statuses SET is_active = ?", vs.IsActive)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
