package controller

import (
	"be-cadidate-votes/models"
	"be-cadidate-votes/utility/binding"
	"be-cadidate-votes/utility/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

// UpdateVoteStatus godoc
// @Summary UpdateVoteStatus
// @Description You can use this API to update/edit vote status in the system.
// @Accept json
// @Tags Private (Bonus)
// @Param Authorization header string true "Bearer ${token}"
// @Param Request body models.VoteStatus true "Request body"
// @Produce json
// @Consumes application/json
// @Success 200 {object} models.GeneralResponse{}
// @Router /api/v1/vote-status [put]
func (ctrl *controller) UpdateVoteStatus(c echo.Context) error {
	ctx := c.Request().Context()
	log := logger.Ctx(ctx)

	// bind and validate body
	req := &models.VoteStatus{}
	if err := binding.GetAndValidateRequestBody(c, log, req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, models.NewErrorWebResponse(err))
	}

	// execute service
	resp, err := ctrl.srv.UpdateVoteStatus(ctx, *req)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, models.NewErrorWebResponse(err))
	}

	return c.JSON(http.StatusOK, resp)
}
