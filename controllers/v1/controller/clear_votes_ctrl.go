package controller

import (
	"be-cadidate-votes/models"
	"be-cadidate-votes/utility/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ClearVotes godoc
// @Summary ClearVotes
// @Description You can use this API to clear all votes in the system.
// @Accept json
// @Tags Private (Vote)
// @Param Authorization header string true "Bearer ${token}"
// @Produce json
// @Consumes application/json
// @Success 200 {object} models.CreateCandidateResponse{}
// @Router /api/v1/votes [delete]
func (ctrl *controller) ClearVotes(c echo.Context) error {
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
