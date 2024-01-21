package binding

import (
	"be-cadidate-votes/models"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"go.uber.org/zap"
)

func BindValidateBody(c echo.Context, request interface{}) error {
	// bind http body and queryparams
	if err := c.Bind(request); err != nil {
		return err
	}
	// validate
	if err := c.Validate(request); err != nil {
		return err
	}
	return nil
}

func GetAndValidateRequestBody(c echo.Context, log *zap.Logger, request interface{}) error {
	if err := BindValidateBody(c, request); err != nil {
		log.Error(err.Error())
		return &models.Err_invalid_param
	}
	return nil
}

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
