package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"golang-auth-service/internal/db"
	"golang-auth-service/internal/handlers"
	"golang-auth-service/internal/repository"
	"golang-auth-service/internal/routes"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load("/Users/anuaradilbek/Desktop/GolangFinal/.env")
	if err != nil {
		log.Println("No .env file found")
	}

	db.InitDB()
	defer db.CloseDB()

	authRepo := &repository.AuthRepository{DB: db.DB}

	authHandler := &handlers.AuthHandler{Repo: authRepo}

	r := mux.NewRouter()

	routes.RegisterRoutes(r, authHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}
