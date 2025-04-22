package http

import (
	"github.com/labstack/echo/v4"
)

func Posts(e *echo.Echo, h *Handler) {
	e.GET("/posts", h.Read)
	e.POST("/posts", h.Create)
	e.PUT("/posts/:id", h.Update)
	e.DELETE("/posts/:id", h.Delete)
}
