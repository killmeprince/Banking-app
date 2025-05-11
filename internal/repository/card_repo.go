package repository

import (
	"banking-app/internal/models"

	"github.com/jmoiron/sqlx"
)

type CardRepo struct{ db *sqlx.DB }

func NewCardRepo(db *sqlx.DB) *CardRepo { return &CardRepo{db} }

func (r *CardRepo) Create(c *models.Card) error {
	return r.db.QueryRowx(
		`INSERT INTO cards 
           (account_id, number_enc, exp_enc, cvv_hash, hmac)
         VALUES ($1, $2, $3, $4, $5)
         RETURNING id, created_at`,
		c.AccountID, c.NumberEnc, c.ExpEnc, c.CVVHash, c.HMAC,
	).StructScan(c)
}

func (r *CardRepo) ByAccount(accID int64) ([]models.Card, error) {
	var xs []models.Card
	err := r.db.Select(&xs, `SELECT * FROM cards WHERE account_id=$1`, accID)
	return xs, err
}
