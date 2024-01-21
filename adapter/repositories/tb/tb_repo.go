package tb

import (
	"be-cadidate-votes/models"
	"context"

	"gorm.io/gorm"
)

type TbRepo interface {
	GetUser(ctx context.Context, filter models.User) (models.User, error)
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
}

func NewTbRepository(db *gorm.DB) TbRepo {
	return &tbRepo{
		db: db,
	}
}

type tbRepo struct {
	db *gorm.DB
}
