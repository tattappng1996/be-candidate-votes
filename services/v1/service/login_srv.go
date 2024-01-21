package service

import (
	"be-cadidate-votes/models"
	"be-cadidate-votes/utility/logger"
	"context"
)

func (srv *service) Login(ctx context.Context, req models.RegisterUserRequest) (models.LoginResponse, error) {
	log := logger.Ctx(ctx)
	response := models.LoginResponse{}

	userInDB, err := srv.repo.TbRepo.GetUser(ctx, models.User{
		Username: req.Username,
	})
	if err != nil {
		log.Error(err.Error())
		return response, &models.Err_backend_system
	}
	if userInDB.ID == 0 {
		return response, &models.Err_data_not_found
	}

	if match := srv.repo.ThirdPartyRepo.CheckPasswordHash(req.Password, userInDB.Password); !match {
		return response, &models.Err_incorrect_data
	}

	accessToken, err := srv.repo.ThirdPartyRepo.CreateJWTToken(srv.cfg.JWT.Secret, srv.cfg.JWT.ClaimExpired, userInDB)
	if err != nil {
		log.Error(err.Error())
		return response, &models.Err_backend_system
	}

	response.Data.AccessToken = accessToken
	response.ResponseStatus = models.SUCCESSResponse

	return response, nil
}
