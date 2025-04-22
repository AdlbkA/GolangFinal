package post

import (
	"blogging-service/pkg/domain"
	"context"
)

type Repository interface {
	CreatePost(ctx context.Context, post domain.PostResponse) (domain.PostResponse, error)
	DeletePost(ctx context.Context, post domain.Post) error
	UpdatePost(ctx context.Context, post domain.Post) (domain.PostResponse, error)
	GetPosts(ctx context.Context) ([]domain.PostResponse, error)
}
