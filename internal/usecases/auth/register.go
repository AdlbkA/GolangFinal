package auth

import (
	"context"
	"golang-auth-service/internal/app/store"
	"golang-auth-service/pkg/auth"
	"golang-auth-service/pkg/domain"
	"golang-auth-service/pkg/reqresp"
)

type RegisterRepository interface {
	RegisterUser(ctx context.Context, username, password, role string) (domain.UserResponse, error)
}

func RegisterUser(ctx context.Context, repo RegisterRepository, req reqresp.RegisterUserRequest) (reqresp.RegisterUserResponse, error) {
	user, err := repo.RegisterUser(ctx, req.Username, req.Password, req.Role)
	if err != nil {
		return reqresp.RegisterUserResponse{}, err
	}

	token, err := auth.CreateToken(user.Username)
	if err != nil {
		return reqresp.RegisterUserResponse{}, err
	}

	return reqresp.RegisterUserResponse{User: user, Token: token}, nil
}

func NewRegisterRepository(st *store.Store) RegisterRepository {
	return &registerRepositoryFacade{st: st}
}

type registerRepositoryFacade struct {
	st *store.Store
}

func (r *registerRepositoryFacade) RegisterUser(ctx context.Context, username, password, role string) (domain.UserResponse, error) {
	return r.st.AuthRepository.CreateUser(ctx, domain.User{Username: username, Password: password, Role: role})
}
