package http

import (
	"blogging-service/pkg/middlewares"
	"github.com/labstack/echo/v4"
)

func Posts(e *echo.Echo, h *Handler) {
	//e.GET("/posts", h.Read)
	//e.POST("/posts", h.Create)
	//e.PUT("/posts/:id", h.Update)
	//e.DELETE("/posts/:id", h.Delete)

	authGroup := e.Group("/posts", middlewares.AuthMiddleware)
	authGroup.GET("", h.Read)
	authGroup.POST("", h.Create)
	authGroup.PUT("/:id", h.Update)
	authGroup.DELETE("/:id", h.Delete)
}
