package reqresp

import "golang-auth-service/pkg/domain"

type RegisterUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type RegisterUserResponse struct {
	User domain.User `json:"user"`
}
