package controller

import (
	"be-cadidate-votes/models"
	"be-cadidate-votes/utility/binding"
	"be-cadidate-votes/utility/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

// UpdateCandidate godoc
// @Summary UpdateCandidate
// @Description You can use this API to update/edit candidate in the system.
// @Accept json
// @Tags Private (Candidate)
// @Param Authorization header string true "Bearer ${token}"
// @Param Request body models.UpdateCandidateRequest true "Request body"
// @Produce json
// @Consumes application/json
// @Success 200 {object} models.CreateCandidateResponse{}
// @Router /api/v1/candidate [put]
func (ctrl *controller) UpdateCandidate(c echo.Context) error {
	ctx := c.Request().Context()
	log := logger.Ctx(ctx)

	// bind and validate body
	req := &models.UpdateCandidateRequest{}
	if err := binding.GetAndValidateRequestBody(c, log, req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, models.NewErrorWebResponse(err))
	}

	// execute service
	resp, err := ctrl.srv.UpdateCandidate(ctx, *req)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, models.NewErrorWebResponse(err))
	}

	return c.JSON(http.StatusOK, resp)
}
