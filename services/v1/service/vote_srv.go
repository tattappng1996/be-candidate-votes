package service

import (
	"be-cadidate-votes/models"
	"be-cadidate-votes/utility/logger"
	"context"
)

func (srv *service) Vote(ctx context.Context, req models.VoteRequest) (models.CreateCandidateResponse, error) {
	log := logger.Ctx(ctx)
	response := models.CreateCandidateResponse{}

	voteInDB, err := srv.repo.TbRepo.GetVote(ctx, req)
	if err != nil {
		log.Error(err.Error())
		return response, &models.Err_backend_system
	}
	if voteInDB.ID > 0 {
		return response, &models.Err_duplicate_data
	}

	candidates, err := srv.repo.TbRepo.ListCandidateWithVote(ctx, models.ListCandidateRequest{
		ID: req.CandidateID,
	})
	if err != nil {
		log.Error(err.Error())
		return response, &models.Err_backend_system
	}
	if len(candidates) == 0 {
		return response, &models.Err_data_not_found
	}

	vote := &models.Vote{
		UserID:      req.UserID,
		CandidateID: req.CandidateID,
	}

	vote, err = srv.repo.TbRepo.CreateVote(ctx, vote)
	if err != nil {
		log.Error(err.Error())
		return response, &models.Err_backend_system
	}

	response.ResponseStatus = models.SUCCESSResponse

	return response, nil
}
