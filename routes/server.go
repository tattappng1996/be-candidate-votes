package routes

import (
	"be-cadidate-votes/appconfig"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func ServerStart(cfg *appconfig.AppConfig, log *zap.Logger) *echo.Echo {

	e := echo.New()

	// Logger
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper: middleware.DefaultSkipper,
		AllowOrigins: []string{
			"*",
		},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodOptions},
	}))

	// routes
	routes(e, cfg, log)

	go func() {
		endPoint := fmt.Sprintf(":%s", cfg.Server.Port)
		if err := e.Start(endPoint); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	return e

}

func ServerShutdown(e *echo.Echo) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
