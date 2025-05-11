package models

import "time"

type Credit struct {
	ID         int64     `db:"id" json:"id"`
	AccountID  int64     `db:"account_id" json:"account_id"`
	Principal  float64   `db:"principal" json:"principal"`
	Rate       float64   `db:"rate" json:"rate"`
	TermMonths int       `db:"term_months" json:"term_months"`
	Margin     float64   `db:"margin" json:"margin"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
}
