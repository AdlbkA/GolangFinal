package start

import (
	"blogging-service/internal/app/config"
	"blogging-service/internal/deliveries/post/http"
	postSrv "blogging-service/internal/service/post"
	"blogging-service/pkg/graceful"
	"context"
	"github.com/labstack/echo/v4"
	"log"
	"time"
)

func HTTP(errs chan<- error, cfg *config.Config, srv postSrv.Service) graceful.Service {
	startType := "http"

	e := echo.New()

	h := http.NewHandler(cfg.HTTP, srv)
	http.Posts(e, h)

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
