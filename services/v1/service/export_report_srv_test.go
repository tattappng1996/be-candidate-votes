package service_test

import (
	"be-cadidate-votes/models"
	"context"
	"fmt"
)

func (t *TestServiceSuite) TestExportReport() {
	t.Run("Should GetVoteStatus Error", func() {
		ctx := context.TODO()

		t.mockTbRepo.EXPECT().
			GetVoteStatus(ctx).
			Return(models.VoteStatus{IsActive: true}, fmt.Errorf("error"))

		expectedResp := models.ReportResponse{}
		expectedErr := &models.Err_backend_system

		actualResp, actualErr := t.srv.ExportReport(ctx)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})
	t.Run("Should ListUser Error", func() {
		ctx := context.TODO()

		t.mockTbRepo.EXPECT().
			GetVoteStatus(ctx).
			Return(models.VoteStatus{IsActive: true}, nil)
		t.mockTbRepo.EXPECT().
			ListUser(ctx, models.User{}).
			Return([]models.User{}, fmt.Errorf("error"))

		expectedResp := models.ReportResponse{
			Data: models.ReportResponseData{
				VoteStatus: models.VoteStatus{
					IsActive: true,
				},
			},
		}
		expectedErr := &models.Err_backend_system

		actualResp, actualErr := t.srv.ExportReport(ctx)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})

	t.Run("Should ListCandidateData Error", func() {
		ctx := context.TODO()

		t.mockTbRepo.EXPECT().
			GetVoteStatus(ctx).
			Return(models.VoteStatus{IsActive: true}, nil)
		t.mockTbRepo.EXPECT().
			ListUser(ctx, models.User{}).
			Return([]models.User{{
				ID: 1,
			}}, nil)
		t.mockTbRepo.EXPECT().
			ListCandidateData(ctx, models.ListCandidateRequest{}).
			Return([]models.ReportCandidateResponse{}, fmt.Errorf("error"))

		expectedResp := models.ReportResponse{
			Data: models.ReportResponseData{
				VoteStatus: models.VoteStatus{
					IsActive: true,
				},
				Users: []models.User{
					{
						ID: 1,
					},
				},
			},
		}
		expectedErr := &models.Err_backend_system

		actualResp, actualErr := t.srv.ExportReport(ctx)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})

	t.Run("Should TestExportReport Success", func() {
		ctx := context.TODO()

		t.mockTbRepo.EXPECT().
			GetVoteStatus(ctx).
			Return(models.VoteStatus{IsActive: true}, nil)
		t.mockTbRepo.EXPECT().
			ListUser(ctx, models.User{}).
			Return([]models.User{{
				ID: 1,
			}}, nil)
		t.mockTbRepo.EXPECT().
			ListCandidateData(ctx, models.ListCandidateRequest{}).
			Return([]models.ReportCandidateResponse{
				{
					ID:     1,
					UserID: 1,
				},
				{
					ID:     1,
					UserID: 2,
				},
				{
					ID:     2,
					UserID: 2,
				},
			}, nil)

		expectedResp := models.ReportResponse{
			ResponseStatus: models.SUCCESSResponse,
			Data: models.ReportResponseData{
				VoteStatus: models.VoteStatus{
					IsActive: true,
				},
				Users: []models.User{
					{
						ID: 1,
					},
				},
				Candidates: []models.ReportCandidateResponse{
					{
						ID:        1,
						VoteCount: 2,
					},
					{
						ID:        2,
						VoteCount: 1,
					},
				},
			},
		}
		var expectedErr error

		actualResp, actualErr := t.srv.ExportReport(ctx)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})
}
