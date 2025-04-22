package app

import (
	"golang-auth-service/internal/app/config"
	"golang-auth-service/internal/app/connections"
	"golang-auth-service/internal/app/start"
	"golang-auth-service/internal/app/store"
	userSrv "golang-auth-service/internal/service/auth"
	"golang-auth-service/pkg/graceful"
	"log"
)

func Run(filenames ...string) {
	cfg, err := config.New(filenames...)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	conns, err := connections.New(cfg)
	if err != nil {
		log.Fatalf("failed to open connections: %v", err)
	}

	st := store.NewRepositoryStore(conns)
	srv := userSrv.NewService(st)

	errs := make(chan error, 50)

	grace := graceful.New(
		start.HTTP(errs, cfg, srv),
	)
	grace.Shutdown(errs, log.Default(), conns)
}
