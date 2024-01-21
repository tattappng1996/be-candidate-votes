package service

import (
	"be-cadidate-votes/adapter/repositories"
	"be-cadidate-votes/appconfig"
	"be-cadidate-votes/models"
	"context"
)

type Service interface {
	RegisterUser(ctx context.Context, req models.RegisterUserRequest) (models.RegisterUserResponse, error)
	Login(ctx context.Context, req models.RegisterUserRequest) (models.LoginResponse, error)

	CreateCandidate(ctx context.Context, req models.CreateCandidateRequest) (models.CreateCandidateResponse, error)
}

func NewService(
	cfg *appconfig.AppConfig,
	repo *repositories.Repo,
) Service {
	return &service{
		cfg,
		repo,
	}
}

type service struct {
	cfg  *appconfig.AppConfig
	repo *repositories.Repo
}
