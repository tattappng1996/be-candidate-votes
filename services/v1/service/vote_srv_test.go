package service_test

import (
	"be-cadidate-votes/models"
	"context"
	"fmt"

	"github.com/golang/mock/gomock"
)

func (t *TestServiceSuite) TestVote() {
	t.Run("Should GetVote Error", func() {
		ctx := context.TODO()
		req := models.VoteRequest{}

		t.mockTbRepo.EXPECT().
			GetVote(ctx, req).
			Return(models.Vote{}, fmt.Errorf("error"))

		expectedResp := models.CreateCandidateResponse{}
		expectedErr := &models.Err_backend_system

		actualResp, actualErr := t.srv.Vote(ctx, req)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})

	t.Run("Should GetVote Found 1", func() {
		ctx := context.TODO()
		req := models.VoteRequest{}

		t.mockTbRepo.EXPECT().
			GetVote(ctx, req).
			Return(models.Vote{ID: 1}, nil)

		expectedResp := models.CreateCandidateResponse{}
		expectedErr := &models.Err_duplicate_data

		actualResp, actualErr := t.srv.Vote(ctx, req)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})

	t.Run("Should ListCandidate Error", func() {
		ctx := context.TODO()
		req := models.VoteRequest{}

		t.mockTbRepo.EXPECT().
			GetVote(ctx, req).
			Return(models.Vote{ID: 0}, nil)
		t.mockTbRepo.EXPECT().
			ListCandidateWithVote(ctx, gomock.Any()).
			Return([]models.CandidateResponse{}, fmt.Errorf("error"))

		expectedResp := models.CreateCandidateResponse{}
		expectedErr := &models.Err_backend_system

		actualResp, actualErr := t.srv.Vote(ctx, req)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})

	t.Run("Should ListCandidate NotFound", func() {
		ctx := context.TODO()
		req := models.VoteRequest{}

		t.mockTbRepo.EXPECT().
			GetVote(ctx, req).
			Return(models.Vote{ID: 0}, nil)
		t.mockTbRepo.EXPECT().
			ListCandidateWithVote(ctx, gomock.Any()).
			Return([]models.CandidateResponse{}, nil)

		expectedResp := models.CreateCandidateResponse{}
		expectedErr := &models.Err_data_not_found

		actualResp, actualErr := t.srv.Vote(ctx, req)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})

	t.Run("Should CreateVote Error", func() {
		ctx := context.TODO()
		req := models.VoteRequest{
			CandidateID: 1,
			UserID:      1,
		}

		t.mockTbRepo.EXPECT().
			GetVote(ctx, req).
			Return(models.Vote{ID: 0}, nil)
		t.mockTbRepo.EXPECT().
			ListCandidateWithVote(ctx, gomock.Any()).
			Return([]models.CandidateResponse{{ID: 1}}, nil)

		vote := &models.Vote{
			CandidateID: 1,
			UserID:      1,
		}
		t.mockTbRepo.EXPECT().
			CreateVote(ctx, vote).
			Return(vote, fmt.Errorf("error"))

		expectedResp := models.CreateCandidateResponse{}
		expectedErr := &models.Err_backend_system

		actualResp, actualErr := t.srv.Vote(ctx, req)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})

	t.Run("Should CreateVote Success", func() {
		ctx := context.TODO()
		req := models.VoteRequest{
			CandidateID: 1,
			UserID:      1,
		}

		t.mockTbRepo.EXPECT().
			GetVote(ctx, req).
			Return(models.Vote{ID: 0}, nil)
		t.mockTbRepo.EXPECT().
			ListCandidateWithVote(ctx, gomock.Any()).
			Return([]models.CandidateResponse{{ID: 1}}, nil)

		vote := &models.Vote{
			CandidateID: 1,
			UserID:      1,
		}
		t.mockTbRepo.EXPECT().
			CreateVote(ctx, vote).
			Return(vote, nil)

		expectedResp := models.CreateCandidateResponse{
			ResponseStatus: models.SUCCESSResponse,
		}
		var expectedErr error

		actualResp, actualErr := t.srv.Vote(ctx, req)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})
}
