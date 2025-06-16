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
		 VALUES ($1, pgp_sym_encrypt($2, $5), pgp_sym_encrypt($3, $5), $4, $5)
		 RETURNING id, created_at`,
		c.AccountID, c.NumberPlain, c.ExpiryPlain, c.CVVHash, c.HMAC,
	).StructScan(c)
}

func (r *CardRepo) ByAccount(accID int64) ([]models.Card, error) {
	var xs []models.Card
	rows, err := r.db.Queryx(
		`SELECT id, account_id, 
		         pgp_sym_decrypt(number_enc, hmac) AS number_plain,
		         pgp_sym_decrypt(exp_enc, hmac) AS expiry_plain,
		         created_at
		   FROM cards
		  WHERE account_id=$1`, accID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var c models.Card
		if err := rows.StructScan(&c); err != nil {
			return nil, err
		}
		xs = append(xs, c)
	}
	return xs, nil
}
func (r *CardRepo) ListByUserID(userID int64) ([]models.Card, error) {
	var cards []models.Card
	err := r.db.Select(&cards, `
		SELECT c.id,
			   c.account_id, 
		       pgp_sym_decrypt(c.number_enc, c.hmac) AS number_plain,
		       pgp_sym_decrypt(c.exp_enc, c.hmac) AS expiry_plain,
		       c.created_at
		  FROM cards c
		  JOIN accounts a ON c.account_id = a.id
		  WHERE a.user_id = $1
	`, userID)
	return cards, err
}
