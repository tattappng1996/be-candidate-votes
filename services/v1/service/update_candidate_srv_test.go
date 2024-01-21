package service_test

import (
	"be-cadidate-votes/models"
	"context"
	"fmt"

	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
)

func (t *TestServiceSuite) TestUpdateCandidate() {
	t.Run("Should UpdateCandidate Error (NotFound)", func() {
		ctx := context.TODO()
		req := models.UpdateCandidateRequest{}

		t.mockTbRepo.EXPECT().
			UpdateCandidate(ctx, gomock.Any()).
			Return(gorm.ErrRecordNotFound)

		expectedResp := models.CreateCandidateResponse{}
		expectedErr := &models.Err_data_not_found

		actualResp, actualErr := t.srv.UpdateCandidate(ctx, req)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})

	t.Run("Should UpdateCandidate Error", func() {
		ctx := context.TODO()
		req := models.UpdateCandidateRequest{}

		t.mockTbRepo.EXPECT().
			UpdateCandidate(ctx, gomock.Any()).
			Return(fmt.Errorf("error"))

		expectedResp := models.CreateCandidateResponse{}
		expectedErr := &models.Err_backend_system

		actualResp, actualErr := t.srv.UpdateCandidate(ctx, req)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})

	t.Run("Should UpdateCandidate Success", func() {
		ctx := context.TODO()
		req := models.UpdateCandidateRequest{}

		t.mockTbRepo.EXPECT().
			UpdateCandidate(ctx, gomock.Any()).
			Return(nil)

		expectedResp := models.CreateCandidateResponse{
			ResponseStatus: models.SUCCESSResponse,
		}
		var expectedErr error

		actualResp, actualErr := t.srv.UpdateCandidate(ctx, req)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})
}
