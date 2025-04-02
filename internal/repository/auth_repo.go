package repository

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"golang-auth-service/internal/db/models"
	"log"
)

type AuthRepository struct {
	DB *sql.DB
}

func (r *AuthRepository) CreateUser(user *models.User) (*models.User, error) {
	check, err := r.DB.Query("Select username from users where username = $1", user.Username)
	if err != nil {
		log.Printf("Failed to add user: %v", err)
		return nil, err
	}
	defer check.Close()

	if check.Next() {
		return nil, errors.New("user already exists")
	}

	_, err = r.DB.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, user.Password)
	if err != nil {
		log.Printf("Failed to add user: %v", err)
		return nil, err
	}

	rows, _ := r.DB.Query("Select username from users where username = $1", user.Username)
	defer rows.Close()
	var username string
	for rows.Next() {
		rows.Scan(&username)

	}

	return user, nil

}
