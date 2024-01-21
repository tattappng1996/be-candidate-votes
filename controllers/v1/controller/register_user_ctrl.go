package controller

import (
	"be-cadidate-votes/models"
	"be-cadidate-votes/utility/binding"
	"be-cadidate-votes/utility/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

// RegisterUser godoc
// @Summary RegisterUser
// @Description You can use this API to create users in the system.
// @Accept json
// @Tags Public (User)
// @Param Request body models.RegisterUserRequest true "Request body"
// @Produce json
// @Consumes application/json
// @Success 200 {object} models.RegisterUserResponse{}
// @Router /api/v1/register [post]
func (ctrl *controller) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()
	log := logger.Ctx(ctx)

	// bind and validate body
	req := &models.RegisterUserRequest{}
	if err := binding.GetAndValidateRequestBody(c, log, req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, models.NewErrorWebResponse(err))
	}

	// execute service
	resp, err := ctrl.srv.RegisterUser(ctx, *req)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, models.NewErrorWebResponse(err))
	}

	return c.JSON(http.StatusOK, resp)
}
