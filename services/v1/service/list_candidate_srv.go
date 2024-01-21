package service

import (
	"be-cadidate-votes/models"
	"be-cadidate-votes/utility/logger"
	"context"
)

func (srv *service) ListCandidate(ctx context.Context, req models.ListCandidateRequest) (models.ListCandidateResponse, error) {
	log := logger.Ctx(ctx)
	response := models.ListCandidateResponse{}

	candidates, err := srv.repo.TbRepo.ListCandidateWithVote(ctx, req)
	if err != nil {
		log.Error(err.Error())
		return response, &models.Err_backend_system
	}

	response.ResponseStatus = models.SUCCESSResponse
	if req.ID > 0 {
		response.Data.Candidates = append(response.Data.Candidates, candidates[0])
		return response, nil
	}

	return response, nil
}
