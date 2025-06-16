package handler

import (
	"banking-app/internal/service"
	"banking-app/pkg/jwt"
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
		AccountID int64  `json:"account_id"`
		Number    string `json:"number"`
		Exp       string `json:"exp"`
		CVV       string `json:"cvv"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}
	card, err := h.svc.Create(req.AccountID, req.Number, req.Exp, req.CVV)
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

func (h *CardHandler) List(w http.ResponseWriter, r *http.Request) {
	userID, err := jwt.GetUserIDFromRequest(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	cards, err := h.svc.ListByUserID(userID)

	if err != nil {
		http.Error(w, "Could not get cards: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cards)
}
