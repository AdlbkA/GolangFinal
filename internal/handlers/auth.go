package handlers

import (
	"encoding/json"
	"golang-auth-service/internal/db/models"
	"golang-auth-service/internal/repository"
	"golang-auth-service/pkg/auth"
	"log"
	"net/http"
)

type AuthHandler struct {
	Repo *repository.AuthRepository
}

type Response struct {
	Username string      `json:"username"`
	Token    interface{} `json:"token"`
	Error    string      `json:"error"`
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := &Response{}

	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp.Error = err.Error()
		log.Fatalf("Error: %v", err)
		return
	}

	res, err := h.Repo.CreateUser(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp.Error = err.Error()
		json.NewEncoder(w).Encode(resp)
		log.Printf("Error: %v", err)
		return
	}

	token, err := auth.CreateToken(res.Username)

	resp.Username = res.Username
	resp.Token = token
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
	log.Println("User created successfully")
}
