package service

import (
	"be-cadidate-votes/models"
	"be-cadidate-votes/utility/logger"
	"context"
)

func (srv *service) CreateCandidate(ctx context.Context, req models.CreateCandidateRequest) (models.CreateCandidateResponse, error) {
	log := logger.Ctx(ctx)
	response := models.CreateCandidateResponse{}

	cInDB, err := srv.repo.TbRepo.GetCandidate(ctx, models.Candidate{
		Name: req.Name,
	})
	if err != nil {
		log.Error(err.Error())
		return response, &models.Err_backend_system
	}
	if cInDB.ID > 0 {
		return response, &models.Err_duplicate_data
	}

	c := &models.Candidate{
		Name:        req.Name,
		Description: req.Description,
	}

	c, err = srv.repo.TbRepo.CreateCandidate(ctx, c)
	if err != nil {
		log.Error(err.Error())
		return response, &models.Err_backend_system
	}

	response.ResponseStatus = models.SUCCESSResponse

	return response, nil
}
