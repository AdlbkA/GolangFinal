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
	//err = middleware.WaitContextCancel(ctx, "RegisterUser", func() error {
	//	var useCaseErr error
	//	resp, useCaseErr = auth.RegisterUser(ctx, auth.NewRegisterRepository(s.st), req)
	//
	//	return useCaseErr
	//})

	resp, err = auth.RegisterUser(ctx, auth.NewRegisterRepository(s.st), req)
	if err != nil {
		log.Printf("auth.RegisterUser err: %v", err)
		return reqresp.RegisterUserResponse{}, err
	}

	return resp, nil
}
