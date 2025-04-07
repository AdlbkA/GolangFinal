package store

import (
	"golang-auth-service/internal/app/connections"
	"golang-auth-service/internal/repository/auth"
	"golang-auth-service/internal/repository/auth/pg"
)

type Store struct {
	AuthRepository auth.Repository
}

func NewRepositoryStore(conns *connections.Connections) *Store {
	st := &Store{
		AuthRepository: pg.NewPostgresRepository(conns.DB),
	}

	return st
}
