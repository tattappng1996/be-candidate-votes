package service

import (
	"be-cadidate-votes/models"
	"be-cadidate-votes/utility/logger"
	"context"
)

func (srv *service) UpdateVoteStatus(ctx context.Context, req models.VoteStatus) (models.GeneralResponse, error) {
	log := logger.Ctx(ctx)
	response := models.GeneralResponse{}

	if err := srv.repo.TbRepo.UpdateVoteStatus(ctx, req); err != nil {
		log.Error(err.Error())
		return response, &models.Err_backend_system
	}

	response.ResponseStatus = models.SUCCESSResponse

	return response, nil
}
