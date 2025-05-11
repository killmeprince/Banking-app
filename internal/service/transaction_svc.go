package service

import (
	"banking-app/internal/models"
	"banking-app/internal/repository"
)

type TransactionService struct {
	repo *repository.TransactionRepo
}

func NewTransactionService(r *repository.TransactionRepo) *TransactionService {
	return &TransactionService{r}
}

func (s *TransactionService) List(accID int64) ([]models.Transaction, error) {
	return s.repo.ByAccount(accID)
}
