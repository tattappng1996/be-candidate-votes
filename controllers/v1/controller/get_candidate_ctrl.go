package controller

import (
	"be-cadidate-votes/models"
	"be-cadidate-votes/utility/logger"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetCandidate godoc
// @Summary GetCandidate
// @Description You can use this API to get candidate by id in the system.
// @Accept json
// @Tags Private
// @Param Authorization header string true "Bearer ${token}"
// @Param id path int true "Candidate ID"
// @Success 200 {object} models.GetCandidateResponse{}
// @Router /api/v1/candidate/{id} [get]
func (ctrl *controller) GetCandidate(c echo.Context) error {
	ctx := c.Request().Context()
	log := logger.Ctx(ctx)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, models.NewErrorWebResponse(err))
	}
	if id == 0 {
		return c.JSON(http.StatusBadRequest, models.Err_invalid_param)
	}

	req := &models.ListCandidateRequest{
		ID: id,
	}

	// execute service
	resp, err := ctrl.srv.ListCandidate(ctx, *req)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, models.NewErrorWebResponse(err))
	}

	return c.JSON(http.StatusOK, models.GetCandidateResponse{
		Data:           resp.Data.Candidates[0],
		ResponseStatus: resp.ResponseStatus,
	})
}
