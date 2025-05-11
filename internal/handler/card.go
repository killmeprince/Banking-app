package handler

import (
	"banking-app/internal/service"
	"encoding/json"
	"net/http"
	"strings"
)

type cardResponse struct {
	ID     int64  `json:"id"`
	Number string `json:"number"`
	Expiry string `json:"expiry"`
}
type CardHandler struct {
	svc *service.CardService
}

func NewCardHandler(s *service.CardService) *CardHandler {
	return &CardHandler{svc: s}
}

func maskNumber(num string) string {
	if len(num) <= 4 {
		return num
	}
	return strings.Repeat("*", len(num)-4) + num[len(num)-4:]
}

func (h *CardHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		AccountID int64 `json:"account_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	card, err := h.svc.Create(req.AccountID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := cardResponse{
		ID:     card.ID,
		Number: maskNumber(card.NumberPlain),
		Expiry: card.ExpiryPlain,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
