package service_test

import (
	"be-cadidate-votes/models"
	"context"
	"fmt"
)

func (t *TestServiceSuite) TestClearVotes() {
	t.Run("Should DeleteVotes Error", func() {
		ctx := context.TODO()

		t.mockTbRepo.EXPECT().
			DeleteVotes(ctx).
			Return(fmt.Errorf("error"))

		expectedResp := models.GeneralResponse{}
		expectedErr := &models.Err_backend_system

		actualResp, actualErr := t.srv.ClearVotes(ctx)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})

	t.Run("Should DeleteVotes Success", func() {
		ctx := context.TODO()

		t.mockTbRepo.EXPECT().
			DeleteVotes(ctx).
			Return(nil)

		expectedResp := models.GeneralResponse{
			ResponseStatus: models.SUCCESSResponse,
		}
		var expectedErr error

		actualResp, actualErr := t.srv.ClearVotes(ctx)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})
}
