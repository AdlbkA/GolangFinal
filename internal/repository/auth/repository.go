package auth

import (
	"context"
	"golang-auth-service/pkg/domain"
)

type Repository interface {
	CreateUser(ctx context.Context, user domain.User) (domain.UserResponse, error)
	//LoginUser(ctx context.Context, user *domain.User) (domain.User, error)
}
