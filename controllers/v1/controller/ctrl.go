package controller

import (
	"be-cadidate-votes/services/v1/service"

	"github.com/labstack/echo/v4"
)

type Controller interface {
	RegisterUser(c echo.Context) error
	Login(c echo.Context) error
}

type controller struct {
	srv service.Service
}

func NewController(
	srv service.Service,
) Controller {
	return &controller{
		srv,
	}
}
