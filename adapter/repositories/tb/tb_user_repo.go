package tb

import (
	"be-cadidate-votes/models"
	"context"
)

func (r *tbRepo) GetUser(ctx context.Context, filter models.User) (models.User, error) {
	query := r.db.WithContext(ctx).Table(`users`)

	if filter.ID != 0 {
		query = query.Where(`id = ?`, filter.ID)
	}
	if filter.Username != "" {
		query = query.Where(`username = ?`, filter.Username)
	}

	u := models.User{}
	if err := query.Scan(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

func (r *tbRepo) ListUser(ctx context.Context, filter models.User) ([]models.User, error) {
	query := r.db.WithContext(ctx).Table(`users`)

	u := []models.User{}
	if err := query.Scan(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

func (r *tbRepo) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	result := r.db.WithContext(ctx).Table(`users`).Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
