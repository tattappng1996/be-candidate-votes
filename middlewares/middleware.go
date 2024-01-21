package middlewares

import (
	"be-cadidate-votes/appconfig"
	"be-cadidate-votes/models"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// AuthMiddleWare ..
func AuthMiddleWare(cfg *appconfig.AppConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get("Authorization")
			if len(token) == 0 {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"status": models.AppError{
						Code:    models.Err_required_token.Code,
						Message: "Token not found in header.",
					},
				})
			}

			splitToken := strings.SplitAfter(token, " ")
			if len(splitToken) == 1 {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"status": models.AppError{
						Code:    models.Err_required_token.Code,
						Message: "Cannot read token in header.",
					},
				})
			}
			token = splitToken[1]

			_, err := verifyToken(cfg.JWT.Secret, token)
			if err != nil {
				log.Println("verifyIDToken error: ", err)
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"status": models.AppError{
						Code:    models.Err_required_token.Code,
						Message: "Token is invalid",
					},
				})
			}

			//c.Set("access_token", token)

			return next(c)
		}
	}
}

func verifyToken(secret string, token string) (map[string]interface{}, error) {
	// sample token is expired.  override time so it parses as valid
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return map[string]interface{}{}, fmt.Errorf("Token is expired")
			}
		}
		return map[string]interface{}{}, err
	}

	if parsedToken == nil {
		return map[string]interface{}{}, fmt.Errorf("Cannot read Token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("Not found key in Token")
	}

	return claims, nil
}
