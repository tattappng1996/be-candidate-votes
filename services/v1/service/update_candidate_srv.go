package service

import (
	"be-cadidate-votes/models"
	"be-cadidate-votes/utility/logger"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (srv *service) UpdateCandidate(ctx context.Context, req models.UpdateCandidateRequest) (models.GeneralResponse, error) {
	log := logger.Ctx(ctx)
	response := models.GeneralResponse{}

	if err := srv.repo.TbRepo.UpdateCandidate(ctx, models.Candidate{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
	}); err != nil {
		log.Error(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response, &models.Err_data_not_found
		}

		return response, &models.Err_backend_system
	}

	response.ResponseStatus = models.SUCCESSResponse

	return response, nil
}
