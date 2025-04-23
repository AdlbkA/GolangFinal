package pg

import (
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
	"golang-auth-service/internal/repository/auth"
	"golang-auth-service/pkg/domain"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type repository struct {
	DB *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) auth.Repository {
	return &repository{DB: db}
}

func (r *repository) CreateUser(ctx context.Context, user domain.User) (domain.UserResponse, error) {
	res := domain.UserResponse{}
	_ = r.DB.Get(&res, "Select username, role from users where username=$1", user.Username)

	if res != (domain.UserResponse{}) {
		return domain.UserResponse{}, errors.New("user already exists")
	}

	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	_, err = r.DB.NamedExec(`INSERT INTO users (username, password) VALUES (:username, :hashedPassword)`,
		map[string]interface{}{
			"username":       user.Username,
			"hashedPassword": string(passwordBytes),
		})
	if err != nil {
		log.Printf("Failed to add user: %v", err)
		return domain.UserResponse{}, err
	}

	err = r.DB.Get(&res, "Select id, username from users where username = $1", user.Username)
	if err != nil {
		log.Println(err)
	}

	return res, nil

}

func (r *repository) LoginUser(ctx context.Context, user domain.User) (domain.UserResponse, error) {
	req := domain.User{}
	err := r.DB.Get(&req, "Select username, password from users where username = $1", user.Username)
	if err != nil {
		log.Println(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(req.Password), []byte(user.Password))
	if err != nil {
		log.Println(err)
		return domain.UserResponse{}, errors.New("invalid username or password")
	}

	res := domain.UserResponse{Username: req.Username}

	return res, nil
}
