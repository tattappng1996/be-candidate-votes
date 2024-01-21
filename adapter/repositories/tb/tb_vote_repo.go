package tb

import (
	"be-cadidate-votes/models"
	"context"

	"gorm.io/gorm"
)

func (r *tbRepo) GetVote(ctx context.Context, filter models.VoteRequest) (models.Vote, error) {
	query := r.db.WithContext(ctx).Table(`votes`)

	if filter.UserID != 0 {
		query = query.Where(`user_id = ?`, filter.UserID)
	}

	v := models.Vote{}
	if err := query.Scan(&v).Error; err != nil {
		return v, err
	}

	return v, nil
}

func (r *tbRepo) CreateVote(ctx context.Context, vote *models.Vote) (*models.Vote, error) {
	result := r.db.WithContext(ctx).Table(`votes`).Create(vote)
	if result.Error != nil {
		return nil, result.Error
	}
	return vote, nil
}

func (r *tbRepo) DeleteVotes(ctx context.Context) error {
	result := r.db.WithContext(ctx).Exec("DELETE FROM votes")
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	if err := r.db.WithContext(ctx).Exec("ALTER SEQUENCE votes_id_seq RESTART WITH 1").Error; err != nil {
		return err
	}

	return nil
}
