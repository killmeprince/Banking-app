package service

import (
	"banking-app/internal/repository"
	"time"
)

type AnalyticsService struct {
	txRepo     *repository.TransactionRepo
	creditRepo *repository.CreditRepo
	accRepo    *repository.AccountRepo
}

func NewAnalyticsService(
	t *repository.TransactionRepo,
	c *repository.CreditRepo,
	a *repository.AccountRepo,
) *AnalyticsService {
	return &AnalyticsService{t, c, a}
}

func (s *AnalyticsService) MonthStats(accID int64) (income, expense float64, err error) {
	txs, err := s.txRepo.ByAccount(accID)
	if err != nil {
		return
	}
	cutoff := time.Now().AddDate(0, -1, 0)
	for _, t := range txs {
		if t.CreatedAt.After(cutoff) {
			if t.Amount > 0 {
				income += t.Amount
			} else {
				expense -= t.Amount
			}
		}
	}
	return
}

func (s *AnalyticsService) CreditLoad(accID int64) (total float64, err error) {
	cs, err := s.creditRepo.ByAccount(accID)
	if err != nil {
		return
	}
	for _, c := range cs {
		total += c.Principal
	}
	return
}

func (s *AnalyticsService) PredictBalance(accID int64, days int) (float64, error) {
	acc, err := s.accRepo.ByID(accID)
	if err != nil {
		return 0, err
	}
	return acc.Balance, nil
}
