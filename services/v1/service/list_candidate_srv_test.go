package service_test

import (
	"be-cadidate-votes/models"
	"context"
	"fmt"

	"github.com/golang/mock/gomock"
)

func (t *TestServiceSuite) TestListCandidate() {
	t.Run("Should ListCandidate Error", func() {
		ctx := context.TODO()
		req := models.ListCandidateRequest{}

		t.mockTbRepo.EXPECT().
			ListCandidateWithVote(ctx, gomock.Any()).
			Return([]models.CandidateResponse{}, fmt.Errorf("error"))

		expectedResp := models.ListCandidateResponse{}
		expectedErr := &models.Err_backend_system

		actualResp, actualErr := t.srv.ListCandidate(ctx, req)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})

	t.Run("Should ListCandidate With ID", func() {
		ctx := context.TODO()
		req := models.ListCandidateRequest{
			ID: 1,
		}

		t.mockTbRepo.EXPECT().
			ListCandidateWithVote(ctx, gomock.Any()).
			Return([]models.CandidateResponse{{ID: 1}}, nil)

		expectedResp := models.ListCandidateResponse{
			Data: models.ListCandidateResponseData{
				Candidates: []models.CandidateResponse{
					{
						ID: 1,
					},
				},
			},
			ResponseStatus: models.SUCCESSResponse,
		}
		var expectedErr error

		actualResp, actualErr := t.srv.ListCandidate(ctx, req)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})
}
