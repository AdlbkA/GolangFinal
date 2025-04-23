package post

import (
	"blogging-service/internal/app/store"
	"blogging-service/internal/usecases/post"
	"blogging-service/pkg/reqresp"
	"context"
	"log"
)

type Service interface {
	CreatePost(ctx context.Context, request reqresp.PostRequest) (reqresp.PostResponse, error)
	UpdatePost(ctx context.Context, request reqresp.PostRequest) (reqresp.PostResponse, error)
	DeletePost(ctx context.Context, request reqresp.PostRequest) (reqresp.PostResponse, error)
	GetPosts(ctx context.Context) (reqresp.PostResponseArray, error)
}

type service struct {
	st *store.Store
}

func NewService(st *store.Store) (srv Service) {
	srv = &service{st: st}
	return srv
}

func (s *service) CreatePost(ctx context.Context, request reqresp.PostRequest) (reqresp.PostResponse, error) {
	resp, err := post.CreatePost(ctx, post.NewPostRepository(s.st), request)
	if err != nil {
		log.Printf("post.CreatePost error: %v", err)
		return reqresp.PostResponse{}, err
	}

	return resp, nil
}

func (s *service) UpdatePost(ctx context.Context, request reqresp.PostRequest) (reqresp.PostResponse, error) {
	resp, err := post.UpdatePost(ctx, post.NewPostRepository(s.st), request)
	if err != nil {
		log.Printf("post.UpdatePost error: %v", err)
		return reqresp.PostResponse{}, err
	}

	return resp, nil
}

func (s *service) DeletePost(ctx context.Context, request reqresp.PostRequest) (reqresp.PostResponse, error) {
	err := post.DeletePost(ctx, post.NewPostRepository(s.st), request.Id)
	if err != nil {
		log.Printf("post.DeletePost error: %v", err)
		return reqresp.PostResponse{}, err
	}

	return reqresp.PostResponse{}, nil
}

func (s *service) GetPosts(ctx context.Context) (reqresp.PostResponseArray, error) {
	resp, err := post.GetPosts(ctx, post.NewPostRepository(s.st))
	if err != nil {
		return reqresp.PostResponseArray{}, err
	}

	return resp, nil

}
