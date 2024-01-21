package controller

import (
	"be-cadidate-votes/models"
	"be-cadidate-votes/utility/binding"
	"be-cadidate-votes/utility/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ListCandidate godoc
// @Summary ListCandidate
// @Description You can use this API to list candidate in the system.
// @Accept json
// @Tags Private
// @Param Authorization header string true "Bearer ${token}"
// @Param Request body models.ListCandidateRequest true "Request body"
// @Produce json
// @Consumes application/json
// @Success 200 {object} models.ListCandidateResponse{}
// @Router /api/v1/list-candidate [post]
func (ctrl *controller) ListCandidate(c echo.Context) error {
	ctx := c.Request().Context()
	log := logger.Ctx(ctx)

	// bind and validate body
	req := &models.ListCandidateRequest{}
	if err := binding.GetAndValidateRequestBody(c, log, req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, models.NewErrorWebResponse(err))
	}
	if req.Limit == 0 {
		req.Limit = models.DefaultLimit
	}
	if req.Page == 0 {
		req.Page = models.DefaultPage
	}

	// execute service
	resp, err := ctrl.srv.ListCandidate(ctx, *req)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, models.NewErrorWebResponse(err))
	}

	return c.JSON(http.StatusOK, resp)
}
