package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"banking-app/internal/service"
)

type TransactionHandler struct{ svc *service.TransactionService }

func NewTransactionHandler(s *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{s}
}

func (h *TransactionHandler) List(w http.ResponseWriter, r *http.Request) {
	accID, _ := strconv.ParseInt(r.URL.Query().Get("account_id"), 10, 64)
	xs, err := h.svc.List(accID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(xs)
}
