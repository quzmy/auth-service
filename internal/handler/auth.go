package handler

import (
	"context"
	"encoding/json"
	"net/http"
)

type AuthServiceInterface interface {
	Register(ctx context.Context, email, password string) error
}

type AuthHandler struct {
	service AuthServiceInterface
}

func NewAuthHandler(service AuthServiceInterface) *AuthHandler {
	return &AuthHandler{service: service}
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if err := h.service.Register(r.Context(), req.Email, req.Password); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("user created"))
}
