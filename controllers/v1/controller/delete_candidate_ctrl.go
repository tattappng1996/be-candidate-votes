package controller

import (
	"be-cadidate-votes/models"
	"be-cadidate-votes/utility/logger"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// DeleteCandidate godoc
// @Summary DeleteCandidate
// @Description You can use this API to delete candidate in the system (Can use when candidate vote is 0).
// @Accept json
// @Tags Private (Candidate)
// @Param Authorization header string true "Bearer ${token}"
// @Param id path int true "Candidate ID"
// @Success 200 {object} models.CreateCandidateResponse{}
// @Router /api/v1/candidate/{id} [delete]
func (ctrl *controller) DeleteCandidate(c echo.Context) error {
	ctx := c.Request().Context()
	log := logger.Ctx(ctx)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, models.NewErrorWebResponse(err))
	}
	if id < 1 {
		return c.JSON(http.StatusBadRequest, models.Err_invalid_param)
	}

	req := &models.ListCandidateRequest{
		ID: id,
	}

	// execute service
	resp, err := ctrl.srv.DeleteCandidate(ctx, *req)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, models.NewErrorWebResponse(err))
	}

	return c.JSON(http.StatusOK, resp)
}
