package http

import (
	"context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang-auth-service/internal/app/config"
	"golang-auth-service/internal/service/auth"
	token "golang-auth-service/pkg/auth"
	"golang-auth-service/pkg/reqresp"
	"net/http"
	"time"
)

type Handler struct {
	srv     auth.Service
	timeout time.Duration
}

func NewHandler(cfg config.HttpConfig, srv auth.Service) *Handler {
	return &Handler{
		srv:     srv,
		timeout: time.Duration(cfg.RequestTimeoutSeconds) * time.Second,
	}
}

func (h *Handler) Register(c echo.Context) error {
	ctx, cancel := h.context(c)

	defer cancel()

	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil
	}

	resp, err := h.srv.RegisterUser(ctx, reqresp.RegisterUserRequest{Username: jsonBody["username"].(string), Password: jsonBody["password"].(string), Role: jsonBody["role"].(string)})
	if err != nil {
		if err.Error() == "user already exists" {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) Login(c echo.Context) error {
	ctx, cancel := h.context(c)

	defer cancel()

	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil
	}

	resp, err := h.srv.LoginUser(ctx, reqresp.LoginUserRequest{Username: jsonBody["username"].(string), Password: jsonBody["password"].(string)})
	if err != nil {
		if err.Error() == "invalid username or password" {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": err.Error(),
			})
		}
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) VerifyToken(c echo.Context) error {
	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil
	}

	resp, err := token.VerifyToken(jsonBody["token"].(string))
	if err != nil {
		if err.Error() == "invalid token" {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": err.Error(),
			})
		}
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) context(c echo.Context) (context.Context, context.CancelFunc) {
	ctx := context.Background()

	ctx = context.WithValue(ctx, "request_id", c.Response().Header().Get(echo.HeaderXRequestID))

	return context.WithTimeout(ctx, h.timeout)
}
