package repository

import (
	"banking-app/internal/models"

	"github.com/jmoiron/sqlx"
)

type CreditRepo struct{ db *sqlx.DB }

func NewCreditRepo(db *sqlx.DB) *CreditRepo { return &CreditRepo{db} }

func (r *CreditRepo) Create(c *models.Credit) error {
	return r.db.QueryRowx(
		`INSERT INTO credits (account_id, principal, rate, term_months, margin)
         VALUES ($1, $2, $3, $4, $5)
         RETURNING id, created_at`,
		c.AccountID, c.Principal, c.Rate, c.TermMonths, c.Margin,
	).StructScan(c)
}

func (r *CreditRepo) ByAccount(accID int64) ([]models.Credit, error) {
	var xs []models.Credit
	err := r.db.Select(&xs, `SELECT * FROM credits WHERE account_id=$1`, accID)
	return xs, err
}
