package service_test

import (
	"be-cadidate-votes/models"
	"context"
	"fmt"

	"github.com/golang/mock/gomock"
)

func (t *TestServiceSuite) TestDeleteCandidate() {
	t.Run("Should ListCandidate Error", func() {
		ctx := context.TODO()
		req := models.ListCandidateRequest{}

		t.mockTbRepo.EXPECT().
			ListCandidateWithVote(ctx, gomock.Any()).
			Return([]models.CandidateResponse{}, fmt.Errorf("error"))

		expectedResp := models.CreateCandidateResponse{}
		expectedErr := &models.Err_backend_system

		actualResp, actualErr := t.srv.DeleteCandidate(ctx, req)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})

	t.Run("Should ListCandidate Error (NotFound)", func() {
		ctx := context.TODO()
		req := models.ListCandidateRequest{}

		t.mockTbRepo.EXPECT().
			ListCandidateWithVote(ctx, gomock.Any()).
			Return([]models.CandidateResponse{}, nil)

		expectedResp := models.CreateCandidateResponse{}
		expectedErr := &models.Err_data_not_found

		actualResp, actualErr := t.srv.DeleteCandidate(ctx, req)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})

	t.Run("Should Get Error (Cannot Delete)", func() {
		ctx := context.TODO()
		req := models.ListCandidateRequest{}

		t.mockTbRepo.EXPECT().
			ListCandidateWithVote(ctx, gomock.Any()).
			Return([]models.CandidateResponse{{ID: 1, VoteCount: 1}}, nil)

		expectedResp := models.CreateCandidateResponse{}
		expectedErr := &models.Err_cannot_delete

		actualResp, actualErr := t.srv.DeleteCandidate(ctx, req)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})

	t.Run("Should UpdateCandidate Error", func() {
		ctx := context.TODO()
		req := models.ListCandidateRequest{}

		t.mockTbRepo.EXPECT().
			ListCandidateWithVote(ctx, gomock.Any()).
			Return([]models.CandidateResponse{{ID: 1, VoteCount: 0}}, nil)
		t.mockTbRepo.EXPECT().
			UpdateCandidate(ctx, gomock.Any()).
			Return(fmt.Errorf("error"))

		expectedResp := models.CreateCandidateResponse{}
		expectedErr := &models.Err_backend_system

		actualResp, actualErr := t.srv.DeleteCandidate(ctx, req)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})

	t.Run("Should DeleteCandidate Success", func() {
		ctx := context.TODO()
		req := models.ListCandidateRequest{}

		t.mockTbRepo.EXPECT().
			ListCandidateWithVote(ctx, gomock.Any()).
			Return([]models.CandidateResponse{{ID: 1, VoteCount: 0}}, nil)
		t.mockTbRepo.EXPECT().
			UpdateCandidate(ctx, gomock.Any()).
			Return(nil)

		expectedResp := models.CreateCandidateResponse{
			ResponseStatus: models.SUCCESSResponse,
		}
		var expectedErr error

		actualResp, actualErr := t.srv.DeleteCandidate(ctx, req)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})
}
