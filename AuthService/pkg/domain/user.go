package domain

type User struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Role     string `json:"role" db:"role"`
}

type UserResponse struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Role     string `json:"role" db:"role"`
}
