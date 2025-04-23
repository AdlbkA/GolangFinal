package reqresp

import "golang-auth-service/pkg/domain"

type RegisterUserRequest struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type RegisterUserResponse struct {
	User  domain.UserResponse `json:"user"`
	Token string              `json:"token"`
}

type LoginUserRequest struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type LoginUserResponse struct {
	User  domain.UserResponse `json:"user"`
	Token string              `json:"token"`
}
