package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"banking-app/internal/service"

	"github.com/gorilla/mux"
)

type CreditHandler struct {
	svc *service.CreditService
}

func NewCreditHandler(s *service.CreditService) *CreditHandler {
	return &CreditHandler{s}
}
func (h *CreditHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		AccountID  int64   `json:"account_id"`
		Principal  float64 `json:"principal"`
		AnnualRate float64 `json:"annual_rate"`
		TermMonths int     `json:"term_months"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	credit, err := h.svc.Create(req.AccountID, req.Principal, req.AnnualRate, req.TermMonths)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(credit)
}

func (h *CreditHandler) Schedule(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "missing credit ID", http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid credit ID", http.StatusBadRequest)
		return
	}

	sched, err := h.svc.GetSchedule(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sched)
}
