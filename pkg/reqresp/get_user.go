package reqresp

import "golang-auth-service/pkg/domain"

type RegisterUserRequest struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Role     string `json:"role" db:"role"`
}

type RegisterUserResponse struct {
	User  domain.UserResponse `json:"user"`
	Token string              `json:"token"`
}
