package service

import (
	"context"
	"fmt"

	"banking-app/internal/models"
	"banking-app/internal/repository"
	"banking-app/pkg/jwt"
)

type AccountService struct {
	repo   *repository.AccountRepo
	txRepo *repository.TransactionRepo
}

func NewAccountService(r *repository.AccountRepo, t *repository.TransactionRepo) *AccountService {
	return &AccountService{r, t}
}

func (s *AccountService) Create(ctx context.Context) (*models.Account, error) {
	sub, ok := ctx.Value(jwt.UserIDKey).(string)
	if !ok {
		return nil, fmt.Errorf("missing user ID")
	}
	var uid int64
	fmt.Sscan(sub, &uid)

	a := &models.Account{
		UserID:  uid,
		Balance: 0,
	}
	if err := s.repo.Create(a); err != nil {
		return nil, err
	}
	return a, nil
}

func (s *AccountService) List(ctx context.Context) ([]models.Account, error) {
	sub := ctx.Value(jwt.UserIDKey).(string)
	var uid int64
	fmt.Sscan(sub, &uid)
	return s.repo.ByUser(uid)
}
func (s *AccountService) Deposit(ctx context.Context, accountID int64, amount float64) (*models.Account, error) {
	sub, ok := ctx.Value(jwt.UserIDKey).(string)
	if !ok {
		return nil, fmt.Errorf("missing user ID")
	}
	var uid int64
	fmt.Sscan(sub, &uid)

	acc, err := s.repo.ByID(accountID)
	if err != nil {
		return nil, err
	}
	if acc.UserID != uid {
		return nil, fmt.Errorf("forbidden")
	}

	if err := s.repo.UpdateBalance(accountID, amount); err != nil {
		return nil, err
	}
	return s.repo.ByID(accountID)
}
func (s *AccountService) Transfer(from, to int64, amount float64) error {
	return s.repo.WithTx(func(tx *repository.AccountRepoTx) error {

		a1, err := tx.ByIDTx(from)
		if err != nil {
			return err
		}
		if a1.Balance < amount {
			return fmt.Errorf("insufficient funds")
		}

		if _, err := tx.ByIDTx(to); err != nil {
			return err
		}

		if err := tx.UpdateBalanceTx(from, -amount); err != nil {
			return err
		}
		if err := tx.UpdateBalanceTx(to, +amount); err != nil {
			return err
		}

		if err := tx.CreateTransactionTx(&models.Transaction{
			AccountID: from,
			Type:      "transfer",
			Amount:    -amount,
		}); err != nil {
			return err
		}
		return tx.CreateTransactionTx(&models.Transaction{
			AccountID: to,
			Type:      "transfer",
			Amount:    amount,
		})
	})
}
