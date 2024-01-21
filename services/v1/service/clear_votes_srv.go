package service

import (
	"be-cadidate-votes/models"
	"be-cadidate-votes/utility/logger"
	"context"
)

func (srv *service) ClearVotes(ctx context.Context) (models.CreateCandidateResponse, error) {
	log := logger.Ctx(ctx)
	response := models.CreateCandidateResponse{}

	if err := srv.repo.TbRepo.DeleteVotes(ctx); err != nil {
		log.Error(err.Error())
		return response, &models.Err_backend_system
	}

	response.ResponseStatus = models.SUCCESSResponse

	return response, nil
}
