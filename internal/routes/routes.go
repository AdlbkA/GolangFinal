package routes

import (
	"github.com/gorilla/mux"
	"golang-auth-service/internal/handlers"
)

func RegisterRoutes(r *mux.Router, authHandler *handlers.AuthHandler) {
	r.HandleFunc("/register", authHandler.Register).Methods("POST")
}
