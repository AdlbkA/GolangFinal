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

	_, err = r.DB.NamedExec(`INSERT INTO users (username, password, role) VALUES (:username, :hashedPassword, :role)`,
		map[string]interface{}{
			"username":       user.Username,
			"hashedPassword": string(passwordBytes),
			"role":           user.Role,
		})
	if err != nil {
		log.Printf("Failed to add user: %v", err)
		return domain.UserResponse{}, err
	}

	err = r.DB.Get(&res, "Select username, role from users where username = $1", user.Username)
	if err != nil {
		log.Println(err)
	}

	return res, nil

}

//func (r *repository) LoginUser(ctx context.Context, user *domain.User) (domain.User, error) {
//	rows, err := r.DB.Query("Select username, password from users where username = $1", user.Username)
//	if err != nil {
//		log.Printf("Failed to get user: %v", err)
//		return domain.User{}, err
//	}
//
//	defer rows.Close()
//
//	if !rows.Next() {
//		return domain.User{}, errors.New("invalid username or password")
//	}
//
//	for rows.Next() {
//		rows.Scan(&user.Username, &user.Password)
//	}
//
//	return user, nil
//
//}
