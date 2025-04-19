package http

import (
	"github.com/labstack/echo/v4"
)

func Register(e *echo.Echo, h *Handler) {
	e.POST("/register", h.Register)
	e.POST("/login", h.Login)
	e.POST("/verify", h.VerifyToken)
}
