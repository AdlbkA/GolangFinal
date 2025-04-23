package middlewares

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

const authServiceURL = "http://golangfinal-auth-service-1:8080/verify" // замени на нужный URL

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header.Get("Authorization")
		if header == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "missing Authorization header")
		}

		parts := strings.Split(header, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid Authorization header format")
		}

		token := parts[1]

		// Prepare JSON body
		body, err := json.Marshal(map[string]string{"token": token})
		if err != nil {
			return err
		}

		resp, err := http.Post(authServiceURL, "application/json", bytes.NewBuffer(body))
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "unable to contact auth service")
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
		}

		// Parse response if it returns user_id or other info
		var data struct {
			UserID string `json:"user_id"`
		}
		respBody, _ := io.ReadAll(resp.Body)
		json.Unmarshal(respBody, &data)

		// Save user ID to context if needed
		c.Set("user_id", data.UserID)

		return next(c)
	}
}
