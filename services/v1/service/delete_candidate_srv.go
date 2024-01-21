package service

import (
	"be-cadidate-votes/models"
	"be-cadidate-votes/utility/logger"
	"context"
	"time"
)

func (srv *service) DeleteCandidate(ctx context.Context, req models.ListCandidateRequest) (models.CreateCandidateResponse, error) {
	log := logger.Ctx(ctx)
	response := models.CreateCandidateResponse{}

	candidates, err := srv.repo.TbRepo.ListCandidateWithVote(ctx, req)
	if err != nil {
		log.Error(err.Error())
		return response, &models.Err_backend_system
	}
	if len(candidates) != 1 {
		return response, &models.Err_data_not_found
	}
	if candidates[0].VoteCount > 0 {
		return response, &models.Err_cannot_delete
	}

	deletedAt := time.Now()

	if err := srv.repo.TbRepo.UpdateCandidate(ctx, models.Candidate{
		ID:        req.ID,
		DeletedAt: &deletedAt,
	}); err != nil {
		log.Error(err.Error())
		return response, &models.Err_backend_system
	}

	response.ResponseStatus = models.SUCCESSResponse

	return response, nil
}
