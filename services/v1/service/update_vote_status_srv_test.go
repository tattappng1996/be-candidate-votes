package service_test

import (
	"be-cadidate-votes/models"
	"context"
	"fmt"

	"github.com/golang/mock/gomock"
)

func (t *TestServiceSuite) TestUpdateVoteStatus() {
	t.Run("Should UpdateVoteStatus Error", func() {
		ctx := context.TODO()
		req := models.VoteStatus{}

		t.mockTbRepo.EXPECT().
			UpdateVoteStatus(ctx, gomock.Any()).
			Return(fmt.Errorf("error"))

		expectedResp := models.GeneralResponse{}
		expectedErr := &models.Err_backend_system

		actualResp, actualErr := t.srv.UpdateVoteStatus(ctx, req)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})

	t.Run("Should UpdateCandidate Success", func() {
		ctx := context.TODO()
		req := models.VoteStatus{}

		t.mockTbRepo.EXPECT().
			UpdateVoteStatus(ctx, gomock.Any()).
			Return(nil)

		expectedResp := models.GeneralResponse{
			ResponseStatus: models.SUCCESSResponse,
		}
		var expectedErr error

		actualResp, actualErr := t.srv.UpdateVoteStatus(ctx, req)

		t.Equal(expectedResp, actualResp, "Response is not equal")
		t.Equal(expectedErr, actualErr, "Error is not equal")
	})
}
