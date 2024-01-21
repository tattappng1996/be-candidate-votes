package service_test

import (
	"be-cadidate-votes/models"
	"context"
	"fmt"

	"github.com/golang/mock/gomock"
)

func (t *TestServiceSuite) TestRegisterUser() {
	t.Run("Should GetUser Error", func() {
		ctx := context.TODO()
		req := models.RegisterUserRequest{}

		t.mockTbRepo.EXPECT().
			GetUser(ctx, gomock.Any()).
			Return(models.User{}, fmt.Errorf("error"))

		expectedResp := models.RegisterUserResponse{}
		expectedErr := &models.Err_backend_system

		actualResp, actualErr := t.srv.RegisterUser(ctx, req)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})

	t.Run("Should GetUser Found User", func() {
		ctx := context.TODO()
		req := models.RegisterUserRequest{}

		t.mockTbRepo.EXPECT().
			GetUser(ctx, gomock.Any()).
			Return(models.User{
				ID: 1,
			}, nil)

		expectedResp := models.RegisterUserResponse{}
		expectedErr := &models.Err_duplicate_user

		actualResp, actualErr := t.srv.RegisterUser(ctx, req)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})

	t.Run("Should HashPassword Error", func() {
		ctx := context.TODO()
		req := models.RegisterUserRequest{}

		t.mockTbRepo.EXPECT().
			GetUser(ctx, gomock.Any()).
			Return(models.User{}, nil)
		t.mockThirdPartyRepo.EXPECT().
			HashPassword(gomock.Any()).
			Return("", fmt.Errorf("error"))

		expectedResp := models.RegisterUserResponse{}
		expectedErr := &models.Err_backend_system

		actualResp, actualErr := t.srv.RegisterUser(ctx, req)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})

	t.Run("Should CreateUser Error", func() {
		ctx := context.TODO()
		req := models.RegisterUserRequest{}

		t.mockTbRepo.EXPECT().
			GetUser(ctx, gomock.Any()).
			Return(models.User{}, nil)
		t.mockThirdPartyRepo.EXPECT().
			HashPassword(gomock.Any()).
			Return("password", nil)

		user := &models.User{
			Password: "password",
			IsActive: true,
		}
		t.mockTbRepo.EXPECT().
			CreateUser(ctx, user).
			Return(user, fmt.Errorf("error"))

		expectedResp := models.RegisterUserResponse{}
		expectedErr := &models.Err_backend_system

		actualResp, actualErr := t.srv.RegisterUser(ctx, req)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})

	t.Run("Should RegisterUser Success", func() {
		ctx := context.TODO()
		req := models.RegisterUserRequest{}

		t.mockTbRepo.EXPECT().
			GetUser(ctx, gomock.Any()).
			Return(models.User{}, nil)
		t.mockThirdPartyRepo.EXPECT().
			HashPassword(gomock.Any()).
			Return("password", nil)

		user := &models.User{
			Password: "password",
			IsActive: true,
		}
		t.mockTbRepo.EXPECT().
			CreateUser(ctx, user).
			Return(user, nil)

		expectedResp := models.RegisterUserResponse{
			ResponseStatus: models.SUCCESSResponse,
		}
		var expectedErr error

		actualResp, actualErr := t.srv.RegisterUser(ctx, req)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})
}
