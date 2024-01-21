package service

import (
	"be-cadidate-votes/adapter/repositories"
	"be-cadidate-votes/appconfig"
	"be-cadidate-votes/models"
	"context"
)

type Service interface {
	RegisterUser(ctx context.Context, req models.RegisterUserRequest) (models.GeneralResponse, error)
	Login(ctx context.Context, req models.RegisterUserRequest) (models.LoginResponse, error)

	CreateCandidate(ctx context.Context, req models.CreateCandidateRequest) (models.GeneralResponse, error)
	UpdateCandidate(ctx context.Context, req models.UpdateCandidateRequest) (models.GeneralResponse, error)
	ListCandidate(ctx context.Context, req models.ListCandidateRequest) (models.ListCandidateResponse, error)
	DeleteCandidate(ctx context.Context, req models.ListCandidateRequest) (models.GeneralResponse, error)
	ClearCandidates(context.Context) (models.GeneralResponse, error)

	Vote(context.Context, models.VoteRequest) (models.GeneralResponse, error)
	ClearVotes(context.Context) (models.GeneralResponse, error)
	UpdateVoteStatus(context.Context, models.VoteStatus) (models.GeneralResponse, error)
	ExportReport(context.Context) (models.ReportResponse, error)
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
