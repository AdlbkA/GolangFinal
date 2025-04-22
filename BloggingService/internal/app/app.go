package app

import (
	"blogging-service/internal/app/config"
	"blogging-service/internal/app/connections"
	"blogging-service/internal/app/start"
	"blogging-service/internal/app/store"
	postSrv "blogging-service/internal/service/post"
	"blogging-service/pkg/graceful"
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
	srv := postSrv.NewService(st)

	errs := make(chan error, 50)

	grace := graceful.New(
		start.HTTP(errs, cfg, srv),
	)
	grace.Shutdown(errs, log.Default(), conns)
}
