package service

import (
	"be-cadidate-votes/models"
	"be-cadidate-votes/utility/logger"
	"context"
)

func (srv *service) ListCandidate(ctx context.Context, req models.ListCandidateRequest) (models.ListCandidateResponse, error) {
	log := logger.Ctx(ctx)
	response := models.ListCandidateResponse{}

	if req.Page > 1 {
		req.Offset = req.Limit * req.Page
	}

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

	count, err := srv.repo.TbRepo.CountCandidate(ctx, req)
	if err != nil {
		log.Error(err.Error())
		return response, &models.Err_backend_system
	}

	response.Data.Candidates = candidates
	response.Data.CandidateCount = count
	response.Data.Limit = req.Limit
	response.Data.Page = req.Page
	response.Data.PageCount = count / req.Limit
	if (count % req.Limit) > 0 {
		response.Data.PageCount = response.Data.PageCount + 1
	}

	return response, nil
}
