package http

import (
	"context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang-auth-service/internal/app/config"
	"golang-auth-service/internal/service/auth"
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

func (h *Handler) context(c echo.Context) (context.Context, context.CancelFunc) {
	ctx := context.Background()

	ctx = context.WithValue(ctx, "request_id", c.Response().Header().Get(echo.HeaderXRequestID))

	return context.WithTimeout(ctx, h.timeout)
}

//type AuthHandler struct {
//	Repo *pg.NewPostgresRepository()
//}
//
//type Response struct {
//	Username string      `json:"username"`
//	Token    interface{} `json:"token"`
//	Error    string      `json:"error"`
//}
//
//func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	resp := &Response{}
//
//	var user domain.User
//
//	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		resp.Error = err.Error()
//		log.Fatalf("Error: %v", err)
//		return
//	}
//
//	res, err := h.Repo.CreateUser(&user)
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		resp.Error = err.Error()
//		json.NewEncoder(w).Encode(resp)
//		log.Printf("Error: %v", err)
//		return
//	}
//
//	token, err := auth.CreateToken(res.Username)
//
//	resp.Username = res.Username
//	resp.Token = token
//	w.WriteHeader(http.StatusCreated)
//	json.NewEncoder(w).Encode(resp)
//	log.Println("User created successfully")
//}
