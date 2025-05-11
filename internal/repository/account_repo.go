package repository

import (
	"banking-app/internal/models"

	"github.com/jmoiron/sqlx"
)

type AccountRepo struct {
	db *sqlx.DB
}

func NewAccountRepo(db *sqlx.DB) *AccountRepo {
	return &AccountRepo{db: db}
}
func (r *AccountRepo) ByUser(userID int64) ([]models.Account, error) {
	var xs []models.Account
	err := r.db.Select(&xs, `SELECT * FROM accounts WHERE user_id=$1`, userID)
	return xs, err
}
func (r *AccountRepo) UpdateBalance(id int64, delta float64) error {
	_, err := r.db.Exec(
		`UPDATE accounts SET balance = balance + $1 WHERE id=$2`,
		delta, id,
	)
	return err
}
func (r *AccountRepo) Create(a *models.Account) error {
	return r.db.QueryRowx(
		`INSERT INTO accounts (user_id, balance)
		 VALUES ($1, $2)
		 RETURNING id, balance, created_at`,
		a.UserID, a.Balance,
	).StructScan(a)
}

func (r *AccountRepo) ByID(id int64) (*models.Account, error) {
	a := new(models.Account)
	err := r.db.Get(a, `SELECT * FROM accounts WHERE id=$1`, id)
	return a, err
}

func (r *AccountRepo) WithTx(fn func(tx *AccountRepoTx) error) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	txRepo := &AccountRepoTx{tx: tx}
	if err := fn(txRepo); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

type AccountRepoTx struct {
	tx *sqlx.Tx
}

func (r *AccountRepoTx) ByIDTx(id int64) (*models.Account, error) {
	a := new(models.Account)
	err := r.tx.Get(a, `SELECT * FROM accounts WHERE id=$1`, id)
	return a, err
}

func (r *AccountRepoTx) UpdateBalanceTx(id int64, delta float64) error {
	_, err := r.tx.Exec(
		`UPDATE accounts SET balance = balance + $1 WHERE id=$2`,
		delta, id,
	)
	return err
}

func (r *AccountRepoTx) CreateTransactionTx(t *models.Transaction) error {
	_, err := r.tx.NamedExec(
		`INSERT INTO transactions (account_id, type, amount, created_at)
		 VALUES (:account_id, :type, :amount, NOW())`,
		t,
	)
	return err
}
