package controller

import (
	"be-cadidate-votes/models"
	"be-cadidate-votes/utility/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ClearCandidates godoc
// @Summary ClearCandidates
// @Description You can use this API to clear all votes in the system. (After clearing all votes user can vote again)
// @Accept json
// @Tags Private (Candidate)
// @Param Authorization header string true "Bearer ${token}"
// @Produce json
// @Consumes application/json
// @Success 200 {object} models.GeneralResponse{}
// @Router /api/v1/candidates [delete]
func (ctrl *controller) ClearCandidates(c echo.Context) error {
	ctx := c.Request().Context()
	log := logger.Ctx(ctx)

	// execute service
	resp, err := ctrl.srv.ClearVotes(ctx)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, models.NewErrorWebResponse(err))
	}

	return c.JSON(http.StatusOK, resp)
}
