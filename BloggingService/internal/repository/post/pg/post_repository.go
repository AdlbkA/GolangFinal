package pg

import (
	"blogging-service/internal/repository/post"
	"blogging-service/pkg/domain"
	"context"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	DB *sqlx.DB
}

func NewPostRepository(db *sqlx.DB) post.Repository {
	return &repository{DB: db}

}

func (r *repository) CreatePost(ctx context.Context, post domain.PostResponse) (domain.PostResponse, error) {
	_, err := r.DB.NamedExec(`INSERT INTO posts VALUES (default, :title, :content, :authorId)`,
		map[string]interface{}{
			"title":    post.Title,
			"content":  post.Content,
			"authorId": post.AuthorId,
		})

	if err != nil {
		return domain.PostResponse{}, err
	}

	return post, nil
}

func (r *repository) DeletePost(ctx context.Context, post domain.Post) error {
	_, err := r.DB.Exec(`DELETE FROM posts WHERE id = $1`, post.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdatePost(ctx context.Context, post domain.Post) (domain.PostResponse, error) {
	_, err := r.DB.Exec(`UPDATE posts SET title = $1, content = $2 WHERE posts.id = $3`, post.Title, post.Content, post.Id)
	if err != nil {
		return domain.PostResponse{}, err
	}

	res := domain.PostResponse{}

	err = r.DB.Get(&res, `SELECT id, title, content, author_id FROM posts WHERE posts.id = $1`, post.Id)
	if err != nil {
		return domain.PostResponse{}, err
	}

	return res, nil
}

func (r *repository) GetPosts(ctx context.Context) ([]domain.PostResponse, error) {
	var posts []domain.PostResponse
	err := r.DB.Get(&posts, `SELECT (id, title, content, author_id) FROM posts`)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
