package reqresp

import "blogging-service/pkg/domain"

type PostRequest struct {
	Id       int    `json:"id" db:"id"`
	Title    string `json:"title" db:"title"`
	Content  string `json:"content" db:"content"`
	AuthorId int    `json:"author_id" db:"author_id"`
}

type PostResponse struct {
	Post domain.PostResponse `json:"post"`
}

type PostResponseArray struct {
	Posts []domain.PostResponse `json:"posts"`
}
