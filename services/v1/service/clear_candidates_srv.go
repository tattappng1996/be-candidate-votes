package service

import (
	"be-cadidate-votes/models"
	"be-cadidate-votes/utility/logger"
	"context"
)

func (srv *service) ClearCandidates(ctx context.Context) (models.GeneralResponse, error) {
	log := logger.Ctx(ctx)
	response := models.GeneralResponse{}

	_, err := srv.ClearVotes(ctx)
	if err != nil {
		log.Error(err.Error())
		return response, &models.Err_backend_system
	}

	if err := srv.repo.TbRepo.DeleteCandidates(ctx); err != nil {
		log.Error(err.Error())
		return response, &models.Err_backend_system
	}

	response.ResponseStatus = models.SUCCESSResponse

	return response, nil
}
