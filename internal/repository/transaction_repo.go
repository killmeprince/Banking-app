package repository

import (
	"banking-app/internal/models"

	"github.com/jmoiron/sqlx"
)

type TransactionRepo struct{ db *sqlx.DB }

func NewTransactionRepo(db *sqlx.DB) *TransactionRepo { return &TransactionRepo{db} }

func (r *TransactionRepo) Create(t *models.Transaction) error {
	_, err := r.db.NamedExec(
		`INSERT INTO transactions (account_id, amount, type) 
      VALUES (:account_id, :amount, :type)`, t)
	return err
}

func (r *TransactionRepo) ByAccount(accID int64) ([]models.Transaction, error) {
	var xs []models.Transaction
	err := r.db.Select(&xs,
		`SELECT * FROM transactions WHERE account_id=$1 ORDER BY created_at DESC`, accID)
	return xs, err
}
