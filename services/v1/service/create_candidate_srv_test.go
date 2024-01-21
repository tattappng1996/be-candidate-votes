package service_test

import (
	"be-cadidate-votes/models"
	"context"
	"fmt"

	"github.com/golang/mock/gomock"
)

func (t *TestServiceSuite) TestCreateCandidate() {
	t.Run("Should GetCandidate Error", func() {
		ctx := context.TODO()
		req := models.CreateCandidateRequest{}

		t.mockTbRepo.EXPECT().
			GetCandidate(ctx, gomock.Any()).
			Return(models.Candidate{}, fmt.Errorf("error"))

		expectedResp := models.CreateCandidateResponse{}
		expectedErr := &models.Err_backend_system

		actualResp, actualErr := t.srv.CreateCandidate(ctx, req)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})

	t.Run("Should GetCandidate And Found", func() {
		ctx := context.TODO()
		req := models.CreateCandidateRequest{}

		t.mockTbRepo.EXPECT().
			GetCandidate(ctx, gomock.Any()).
			Return(models.Candidate{
				ID: 1,
			}, nil)

		expectedResp := models.CreateCandidateResponse{}
		expectedErr := &models.Err_duplicate_data

		actualResp, actualErr := t.srv.CreateCandidate(ctx, req)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})

	t.Run("Should CreateCandidate Error", func() {
		ctx := context.TODO()
		req := models.CreateCandidateRequest{}

		t.mockTbRepo.EXPECT().
			GetCandidate(ctx, gomock.Any()).
			Return(models.Candidate{}, nil)

		c := &models.Candidate{}
		t.mockTbRepo.EXPECT().
			CreateCandidate(ctx, gomock.Any()).
			Return(c, fmt.Errorf("error"))

		expectedResp := models.CreateCandidateResponse{}
		expectedErr := &models.Err_backend_system

		actualResp, actualErr := t.srv.CreateCandidate(ctx, req)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})

	t.Run("Should CreateCandidate Success", func() {
		ctx := context.TODO()
		req := models.CreateCandidateRequest{}

		t.mockTbRepo.EXPECT().
			GetCandidate(ctx, gomock.Any()).
			Return(models.Candidate{}, nil)

		c := &models.Candidate{}
		t.mockTbRepo.EXPECT().
			CreateCandidate(ctx, gomock.Any()).
			Return(c, nil)

		expectedResp := models.CreateCandidateResponse{
			ResponseStatus: models.SUCCESSResponse,
		}
		var expectedErr error

		actualResp, actualErr := t.srv.CreateCandidate(ctx, req)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})
}
