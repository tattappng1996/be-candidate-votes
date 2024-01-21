package controller

import (
	"be-cadidate-votes/models"
	"be-cadidate-votes/utility/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ExportReport godoc
// @Summary ExportReport
// @Description You can use this API to get report in the system.
// @Accept json
// @Tags Private (Bonus)
// @Param Authorization header string true "Bearer ${token}"
// @Success 200 {object} models.ReportResponse{}
// @Router /api/v1/report [get]
func (ctrl *controller) ExportReport(c echo.Context) error {
	ctx := c.Request().Context()
	log := logger.Ctx(ctx)

	// execute service
	resp, err := ctrl.srv.ExportReport(ctx)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, models.NewErrorWebResponse(err))
	}

	return c.JSON(http.StatusOK, resp)
}
