package handler

import (
	"encoding/json"
	"net/http"

	"banking-app/internal/service"
)

type AccountHandler struct{ svc *service.AccountService }

func NewAccountHandler(s *service.AccountService) *AccountHandler {
	return &AccountHandler{s}
}
func (h *AccountHandler) Deposit(w http.ResponseWriter, r *http.Request) {
	var req struct {
		AccountID int64   `json:"account_id"`
		Amount    float64 `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	acc, err := h.svc.Deposit(r.Context(), req.AccountID, req.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(acc)
}
func (h *AccountHandler) Create(w http.ResponseWriter, r *http.Request) {
	acc, err := h.svc.Create(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(acc)
}

func (h *AccountHandler) List(w http.ResponseWriter, r *http.Request) {
	xs, err := h.svc.List(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(xs)
}

func (h *AccountHandler) Transfer(w http.ResponseWriter, r *http.Request) {
	var req struct {
		FromID int64   `json:"from_account_id"`
		ToID   int64   `json:"to_account_id"`
		Amount float64 `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	if err := h.svc.Transfer(req.FromID, req.ToID, req.Amount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "transferred"})
}
