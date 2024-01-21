package tb

import (
	"context"

	"gorm.io/gorm"
)

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
