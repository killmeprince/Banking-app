package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"banking-app/internal/service"
)

type AnalyticsHandler struct{ svc *service.AnalyticsService }

func NewAnalyticsHandler(s *service.AnalyticsService) *AnalyticsHandler {
	return &AnalyticsHandler{s}
}

func (h *AnalyticsHandler) MonthStats(w http.ResponseWriter, r *http.Request) {
	accID, _ := strconv.ParseInt(r.URL.Query().Get("account_id"), 10, 64)
	income, expense, err := h.svc.MonthStats(accID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(map[string]float64{"income": income, "expense": expense})
}

func (h *AnalyticsHandler) CreditLoad(w http.ResponseWriter, r *http.Request) {
	accID, _ := strconv.ParseInt(r.URL.Query().Get("account_id"), 10, 64)
	total, err := h.svc.CreditLoad(accID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(map[string]float64{"credit_load": total})
}

func (h *AnalyticsHandler) Predict(w http.ResponseWriter, r *http.Request) {
	accID, _ := strconv.ParseInt(r.URL.Query().Get("account_id"), 10, 64)
	days, _ := strconv.Atoi(r.URL.Query().Get("days"))
	bal, err := h.svc.PredictBalance(accID, days)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(map[string]float64{"predicted_balance": bal})
}
