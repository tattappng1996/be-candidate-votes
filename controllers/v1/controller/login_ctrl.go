package controller

import (
	"be-cadidate-votes/models"
	"be-cadidate-votes/utility/binding"
	"be-cadidate-votes/utility/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Login godoc
// @Summary Login
// @Description Users need to use this API to log in and grant access token.
// @Accept json
// @Tags Public (User)
// @Param Request body models.RegisterUserRequest true "Request body"
// @Produce json
// @Consumes application/json
// @Success 200 {object} models.LoginResponse{}
// @Router /api/v1/login [post]
func (ctrl *controller) Login(c echo.Context) error {
	ctx := c.Request().Context()
	log := logger.Ctx(ctx)

	// bind and validate body
	req := &models.RegisterUserRequest{}
	if err := binding.GetAndValidateRequestBody(c, log, req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, models.NewErrorWebResponse(err))
	}

	// execute service
	resp, err := ctrl.srv.Login(ctx, *req)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, models.NewErrorWebResponse(err))
	}

	return c.JSON(http.StatusOK, resp)
}
