package auth

import (
	"context"
	"golang-auth-service/internal/app/store"
	"golang-auth-service/internal/usecases/auth"
	"golang-auth-service/pkg/reqresp"
	"log"
)

type Service interface {
	RegisterUser(ctx context.Context, request reqresp.RegisterUserRequest) (reqresp.RegisterUserResponse, error)
	LoginUser(ctx context.Context, request reqresp.LoginUserRequest) (reqresp.LoginUserResponse, error)
}

type service struct {
	st *store.Store
}

func NewService(st *store.Store) (srv Service) {
	srv = &service{st: st}
	return srv
}

func (s *service) RegisterUser(
	ctx context.Context,
	req reqresp.RegisterUserRequest,
) (resp reqresp.RegisterUserResponse, err error) {

	resp, err = auth.RegisterUser(ctx, auth.NewRegisterRepository(s.st), req)
	if err != nil {
		log.Printf("auth.RegisterUser err: %v", err)
		return reqresp.RegisterUserResponse{}, err
	}

	return resp, nil
}

func (s *service) LoginUser(
	ctx context.Context,
	request reqresp.LoginUserRequest,
) (reqresp.LoginUserResponse, error) {
	resp, err := auth.LoginUser(ctx, auth.NewRegisterRepository(s.st), request)
	if err != nil {
		log.Printf("auth.LoginUser err: %v", err)
		return reqresp.LoginUserResponse{}, err
	}

	return resp, nil
}
