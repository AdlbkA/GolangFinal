package domain

type Post struct {
	Id       int    `json:"id" db:"id"`
	Title    string `json:"title" db:"title"`
	Content  string `json:"content" db:"content"`
	AuthorId int    `json:"author_id" db:"author_id"`
}

type PostResponse struct {
	ID       int    `json:"id" db:"id"`
	Title    string `json:"title" db:"title"`
	Content  string `json:"content" db:"content"`
	AuthorId int    `json:"author_id" db:"author_id"`
}
