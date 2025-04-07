package start

import (
	"context"
	"github.com/labstack/echo/v4"
	"golang-auth-service/internal/app/config"
	"golang-auth-service/internal/deliveries/auth/http"
	userSrv "golang-auth-service/internal/service/auth"
	"golang-auth-service/pkg/graceful"
	"log"
	"time"
)

func HTTP(errs chan<- error, cfg *config.Config, srv userSrv.Service) graceful.Service {
	startType := "http"

	e := echo.New()

	h := http.NewHandler(cfg.HTTP, srv)
	http.Register(e, h)

	go func() {
		errs <- e.Start(cfg.HTTP.Addr)
	}()

	return graceful.NewService(startType, func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := e.Shutdown(ctx); err != nil {
			log.Printf("http shutdown error: %v", err)
		}
	})
}
