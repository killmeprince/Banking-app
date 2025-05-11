package handler

import (
	"encoding/json"
	"net/http"

	"banking-app/internal/service"
)

type UserHandler struct{ svc *service.UserService }

func NewUserHandler(s *service.UserService) *UserHandler { return &UserHandler{s} }

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req struct{ Email, Password string }
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	if err := h.svc.Register(req.Email, req.Password); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "created"})
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct{ Email, Password string }
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	tok, err := h.svc.Login(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": tok})
}
