package service

import (
	"be-cadidate-votes/models"
	"be-cadidate-votes/utility/logger"
	"context"
)

func (srv *service) CreateCandidate(ctx context.Context, req models.CreateCandidateRequest) (models.GeneralResponse, error) {
	log := logger.Ctx(ctx)
	response := models.GeneralResponse{}

	_, err := srv.repo.TbRepo.CreateCandidate(ctx, &models.Candidate{
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		log.Error(err.Error())
		return response, &models.Err_backend_system
	}

	response.ResponseStatus = models.SUCCESSResponse

	return response, nil
}
