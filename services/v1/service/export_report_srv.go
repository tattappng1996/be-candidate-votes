package service

import (
	"be-cadidate-votes/models"
	"be-cadidate-votes/utility/logger"
	"context"
)

func (srv *service) ExportReport(ctx context.Context) (models.ReportResponse, error) {
	log := logger.Ctx(ctx)
	response := models.ReportResponse{}

	vs, err := srv.repo.TbRepo.GetVoteStatus(ctx)
	if err != nil {
		log.Error(err.Error())
		return response, &models.Err_backend_system
	}
	response.Data.VoteStatus = vs

	users, err := srv.repo.TbRepo.ListUser(ctx, models.User{})
	if err != nil {
		log.Error(err.Error())
		return response, &models.Err_backend_system
	}
	if len(users) > 0 {
		userIndexMap := map[int]int{}
		for i, user := range users {
			response.Data.Users = append(response.Data.Users, user)
			userIndexMap[user.ID] = i
		}

		listData, err := srv.repo.TbRepo.ListCandidateData(ctx, models.ListCandidateRequest{})
		if err != nil {
			log.Error(err.Error())
			return response, &models.Err_backend_system
		}
		if len(listData) > 0 {
			candidateIndexMap := map[int]int{}
			for _, data := range listData {
				if data.UserID > 0 {
					data.VoteCount = 1
					response.Data.Users[userIndexMap[data.UserID]].VotedTo = data.Name
				}

				if _, dup := candidateIndexMap[data.ID]; !dup {
					response.Data.Candidates = append(response.Data.Candidates, models.ReportCandidateResponse{
						ID:          data.ID,
						Name:        data.Name,
						Description: data.Description,
						VoteCount:   data.VoteCount,
					})

					candidateIndexMap[data.ID] = len(response.Data.Candidates) - 1
				} else {
					response.Data.Candidates[candidateIndexMap[data.ID]].VoteCount += data.VoteCount
				}
			}
		}
	}

	response.ResponseStatus = models.SUCCESSResponse

	return response, nil
}
