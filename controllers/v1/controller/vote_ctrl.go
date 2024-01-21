package controller

import (
	"be-cadidate-votes/models"
	"be-cadidate-votes/utility/binding"
	"be-cadidate-votes/utility/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Vote godoc
// @Summary Vote
// @Description You can use this API to vote for candidates in the system. (User can only have 1 vote)
// @Accept json
// @Tags Private (Vote)
// @Param Authorization header string true "Bearer ${token}"
// @Param Request body models.VoteRequest true "Request body"
// @Produce json
// @Consumes application/json
// @Success 200 {object} models.CreateCandidateResponse{}
// @Router /api/v1/vote [post]
func (ctrl *controller) Vote(c echo.Context) error {
	ctx := c.Request().Context()
	log := logger.Ctx(ctx)

	// bind and validate body
	req := &models.VoteRequest{}
	if err := binding.GetAndValidateRequestBody(c, log, req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, models.NewErrorWebResponse(err))
	}
	req.UserID = c.Get("user_id").(int)

	// execute service
	resp, err := ctrl.srv.Vote(ctx, *req)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, models.NewErrorWebResponse(err))
	}

	return c.JSON(http.StatusOK, resp)
}
