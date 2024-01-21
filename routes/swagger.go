package routes

import (
	"be-cadidate-votes/appconfig"
	"fmt"
	"net/http"

	echoDocsV1 "be-cadidate-votes/docs/v1"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func InitSwagger(e *echo.Echo, cfg *appconfig.AppConfig) {
	if cfg.Server.SwaggerEnable {
		echoDocsV1.SwaggerInfo.Title = "be-candidate-votes-service"
		echoDocsV1.SwaggerInfo.Description = "This is a be-candidate-votes-service server."
		echoDocsV1.SwaggerInfo.Version = "1.0"
		echoDocsV1.SwaggerInfo.Host = cfg.Server.Url
		echoDocsV1.SwaggerInfo.Schemes = []string{"http", "https"}

		if cfg.Server.Url == "" {
			echoDocsV1.SwaggerInfo.Host = fmt.Sprintf(`%s:%s`, "localhost", cfg.Server.Port)
		}

		e.GET("/api/v1/swagger/*", echoSwagger.WrapHandler)
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			Skipper: middleware.DefaultSkipper,
			AllowOrigins: []string{
				fmt.Sprintf(`http://%s`, echoDocsV1.SwaggerInfo.Host),
				fmt.Sprintf(`https://%s`, echoDocsV1.SwaggerInfo.Host),
			},
			AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodOptions},
		}))
	}
}
