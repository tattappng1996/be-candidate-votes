package routes

import (
	"be-cadidate-votes/adapter/repositories"
	"be-cadidate-votes/appconfig"
	"be-cadidate-votes/controllers/v1/controller"
	"be-cadidate-votes/middlewares"
	"be-cadidate-votes/services/v1/service"
	"be-cadidate-votes/utility/binding"
	"be-cadidate-votes/utility/sql/postgres"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func routes(e *echo.Echo, cfg *appconfig.AppConfig, log *zap.Logger) {
	db, err := postgres.NewPostgresqlDatabase(cfg.Database.DSN)
	if err != nil {
		log.Sugar().Fatalf("fatal error connect postgresql: %v", err)
	}

	// provide repositories
	repo := repositories.NewRepository(db)

	// provide services
	srv := service.NewService(cfg, repo)

	// provide controller
	ctrl := controller.NewController(srv)

	// Custom Validator
	e.Validator = &binding.CustomValidator{Validator: validator.New()}

	e.GET("/", healthCheck)

	// group
	v1 := e.Group("/api/v1")
	InitSwagger(e, cfg)

	// Public
	v1.POST("/register", ctrl.RegisterUser)
	v1.POST("/login", ctrl.Login)

	// Private
	v1.Use(middlewares.AuthMiddleWare(cfg))
	v1.POST("/candidate", ctrl.CreateCandidate)
	v1.GET("/candidate/:id", ctrl.GetCandidate)
	v1.PUT("/candidate", ctrl.UpdateCandidate)
	v1.POST("/list-candidate", ctrl.ListCandidate)
	v1.DELETE("/candidate/:id", ctrl.DeleteCandidate)
	//v1.DELETE("/candidates", ctrl.ClearCandidates)
	v1.POST("/vote", ctrl.Vote)
	v1.DELETE("/votes", ctrl.ClearVotes)
	// bonus
	//v1.GET("/report", ctrl.ExportReport)
	//v1.POST("/vote-state", ctrl.OpenOrCloseVote)
}

// HealthCheck godoc
// @Summary HealthCheck
// @Description Just HealthCheck
// @Tags Generals
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{} "response"
// @Router / [get]
func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"service":   "BE-CANDIDATE-VOTES-SERVICES.",
		"timestamp": time.Now().Format(time.RFC3339),
		"success":   true,
	})
}
