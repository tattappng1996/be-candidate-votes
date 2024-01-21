package service_test

import (
	"be-cadidate-votes/models"
	"context"
	"fmt"
)

func (t *TestServiceSuite) TestClearCandidates() {
	t.Run("Should DeleteVotes Error", func() {
		ctx := context.TODO()

		t.mockTbRepo.EXPECT().
			DeleteVotes(ctx).
			Return(fmt.Errorf("error"))

		expectedResp := models.GeneralResponse{}
		expectedErr := &models.Err_backend_system

		actualResp, actualErr := t.srv.ClearCandidates(ctx)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})
	t.Run("Should DeleteCandidates Error", func() {
		ctx := context.TODO()

		t.mockTbRepo.EXPECT().
			DeleteVotes(ctx).
			Return(nil)
		t.mockTbRepo.EXPECT().
			DeleteCandidates(ctx).
			Return(fmt.Errorf("error"))

		expectedResp := models.GeneralResponse{}
		expectedErr := &models.Err_backend_system

		actualResp, actualErr := t.srv.ClearCandidates(ctx)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})
	t.Run("Should TestClearCandidates Success", func() {
		ctx := context.TODO()

		t.mockTbRepo.EXPECT().
			DeleteVotes(ctx).
			Return(nil)
		t.mockTbRepo.EXPECT().
			DeleteCandidates(ctx).
			Return(nil)

		expectedResp := models.GeneralResponse{
			ResponseStatus: models.SUCCESSResponse,
		}
		var expectedErr error

		actualResp, actualErr := t.srv.ClearCandidates(ctx)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})
}
