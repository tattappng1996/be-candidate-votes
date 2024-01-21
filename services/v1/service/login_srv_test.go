package service_test

import (
	"be-cadidate-votes/models"
	"context"
	"fmt"

	"github.com/golang/mock/gomock"
)

func (t *TestServiceSuite) TestLogin() {
	t.Run("Should GetUser Error", func() {
		ctx := context.TODO()
		req := models.RegisterUserRequest{}

		t.mockTbRepo.EXPECT().
			GetUser(ctx, gomock.Any()).
			Return(models.User{}, fmt.Errorf("error"))

		expectedResp := models.LoginResponse{}
		expectedErr := &models.Err_backend_system

		actualResp, actualErr := t.srv.Login(ctx, req)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})

	t.Run("Should GetUser ID Equal 0", func() {
		ctx := context.TODO()
		req := models.RegisterUserRequest{}

		t.mockTbRepo.EXPECT().
			GetUser(ctx, gomock.Any()).
			Return(models.User{}, nil)

		expectedResp := models.LoginResponse{}
		expectedErr := &models.Err_data_not_found

		actualResp, actualErr := t.srv.Login(ctx, req)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})

	t.Run("Should CheckPasswordHash Not Match", func() {
		ctx := context.TODO()
		req := models.RegisterUserRequest{}

		t.mockTbRepo.EXPECT().
			GetUser(ctx, gomock.Any()).
			Return(models.User{
				ID: 1,
			}, nil)

		t.mockThirdPartyRepo.EXPECT().
			CheckPasswordHash(gomock.Any(), gomock.Any()).
			Return(false)

		expectedResp := models.LoginResponse{}
		expectedErr := &models.Err_incorrect_data

		actualResp, actualErr := t.srv.Login(ctx, req)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})

	t.Run("Should CreateJWTToken Error", func() {
		ctx := context.TODO()
		req := models.RegisterUserRequest{}

		t.mockTbRepo.EXPECT().
			GetUser(ctx, gomock.Any()).
			Return(models.User{
				ID: 1,
			}, nil)
		t.mockThirdPartyRepo.EXPECT().
			CheckPasswordHash(gomock.Any(), gomock.Any()).
			Return(true)
		t.mockThirdPartyRepo.EXPECT().
			CreateJWTToken(gomock.Any(), gomock.Any(), gomock.Any()).
			Return("", fmt.Errorf("error"))

		expectedResp := models.LoginResponse{}
		expectedErr := &models.Err_backend_system

		actualResp, actualErr := t.srv.Login(ctx, req)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})

	t.Run("Should Login Success", func() {
		ctx := context.TODO()
		req := models.RegisterUserRequest{}

		t.mockTbRepo.EXPECT().
			GetUser(ctx, gomock.Any()).
			Return(models.User{
				ID: 1,
			}, nil)
		t.mockThirdPartyRepo.EXPECT().
			CheckPasswordHash(gomock.Any(), gomock.Any()).
			Return(true)
		t.mockThirdPartyRepo.EXPECT().
			CreateJWTToken(gomock.Any(), gomock.Any(), gomock.Any()).
			Return("token", nil)

		expectedResp := models.LoginResponse{
			Data: models.LoginResponseData{
				AccessToken: "token",
			},
			ResponseStatus: models.SUCCESSResponse,
		}
		var expectedErr error

		actualResp, actualErr := t.srv.Login(ctx, req)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})
}
