package post

import (
	"blogging-service/internal/app/store"
	"blogging-service/pkg/domain"
	"blogging-service/pkg/reqresp"
	"context"
)

type BlogRepository interface {
	CreatePost(ctx context.Context, title, content string, authorId int) (domain.PostResponse, error)
	UpdatePost(ctx context.Context, id int, title, content string, authorId int) (domain.PostResponse, error)
	DeletePost(ctx context.Context, id int) error
	GetPosts(ctx context.Context) ([]domain.PostResponse, error)
}

func CreatePost(ctx context.Context, repo BlogRepository, req reqresp.PostRequest) (reqresp.PostResponse, error) {
	post, err := repo.CreatePost(ctx, req.Title, req.Content, req.AuthorId)
	if err != nil {
		return reqresp.PostResponse{}, err
	}

	return reqresp.PostResponse{Post: post}, nil
}

func UpdatePost(ctx context.Context, repo BlogRepository, req reqresp.PostRequest) (reqresp.PostResponse, error) {
	post, err := repo.UpdatePost(ctx, req.Id, req.Title, req.Content, req.AuthorId)
	if err != nil {
		return reqresp.PostResponse{}, err
	}

	return reqresp.PostResponse{Post: post}, nil
}

func DeletePost(ctx context.Context, repo BlogRepository, id int) error {
	err := repo.DeletePost(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func GetPosts(ctx context.Context, repo BlogRepository) (reqresp.PostResponseArray, error) {
	posts, err := repo.GetPosts(ctx)
	if err != nil {
		return reqresp.PostResponseArray{}, err
	}

	return reqresp.PostResponseArray{Posts: posts}, nil
}

func NewPostRepository(st *store.Store) BlogRepository {
	return &postRepositoryFacade{st: st}
}

type postRepositoryFacade struct {
	st *store.Store
}

func (p *postRepositoryFacade) CreatePost(ctx context.Context, title, content string, authorId int) (domain.PostResponse, error) {
	return p.st.PostRepository.CreatePost(ctx, domain.PostResponse{Title: title, Content: content, AuthorId: authorId})
}

func (p *postRepositoryFacade) UpdatePost(ctx context.Context, id int, title, content string, authorId int) (domain.PostResponse, error) {
	return p.st.PostRepository.UpdatePost(ctx, domain.Post{
		Id:       id,
		Title:    title,
		Content:  content,
		AuthorId: authorId,
	})

}

func (p *postRepositoryFacade) DeletePost(ctx context.Context, id int) error {
	return p.st.PostRepository.DeletePost(ctx, domain.Post{Id: id})
}

func (p *postRepositoryFacade) GetPosts(ctx context.Context) ([]domain.PostResponse, error) {
	return p.st.PostRepository.GetPosts(ctx)
}
