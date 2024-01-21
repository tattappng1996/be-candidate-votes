package controller

import (
	"be-cadidate-votes/services/v1/service"

	"github.com/labstack/echo/v4"
)

type Controller interface {
	// public
	RegisterUser(c echo.Context) error
	Login(c echo.Context) error

	// private
	CreateCandidate(c echo.Context) error
	GetCandidate(c echo.Context) error
	UpdateCandidate(c echo.Context) error
	ListCandidate(c echo.Context) error
	DeleteCandidate(c echo.Context) error
	ClearCandidates(c echo.Context) error
	Vote(c echo.Context) error
	ClearVotes(c echo.Context) error
	ExportReport(c echo.Context) error
	UpdateVoteStatus(c echo.Context) error
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
