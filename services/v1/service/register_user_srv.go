package service

import (
	"be-cadidate-votes/models"
	"be-cadidate-votes/utility/logger"
	"context"
)

func (srv *service) RegisterUser(ctx context.Context, req models.RegisterUserRequest) (models.GeneralResponse, error) {
	log := logger.Ctx(ctx)
	response := models.GeneralResponse{}

	userInDB, err := srv.repo.TbRepo.GetUser(ctx, models.User{
		Username: req.Username,
	})
	if err != nil {
		log.Error(err.Error())
		return response, &models.Err_backend_system
	}
	if userInDB.ID > 0 {
		return response, &models.Err_duplicate_data
	}

	hashPassword, err := srv.repo.ThirdPartyRepo.HashPassword(req.Password)
	if err != nil {
		log.Error(err.Error())
		return response, &models.Err_backend_system
	}
	user := &models.User{
		Username: req.Username,
		Password: hashPassword,
		IsActive: true,
	}

	user, err = srv.repo.TbRepo.CreateUser(ctx, user)
	if err != nil {
		log.Error(err.Error())
		return response, &models.Err_backend_system
	}

	response.ResponseStatus = models.SUCCESSResponse

	return response, nil
}
