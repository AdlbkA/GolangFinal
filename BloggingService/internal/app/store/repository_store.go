package store

import (
	"blogging-service/internal/app/connections"
	"blogging-service/internal/repository/post"
	"blogging-service/internal/repository/post/pg"
)

type Store struct {
	PostRepository post.Repository
}

func NewRepositoryStore(conns *connections.Connections) *Store {
	st := &Store{
		PostRepository: pg.NewPostRepository(conns.DB),
	}

	return st
}
