package models

import "time"

type Card struct {
	ID          int64     `db:"id" json:"id"`
	AccountID   int64     `db:"account_id" json:"account_id"`
	NumberEnc   []byte    `db:"number_enc" json:"-"`
	ExpEnc      []byte    `db:"exp_enc" json:"-"`
	CVVHash     string    `db:"cvv_hash" json:"-"`
	HMAC        string    `db:"hmac" json:"-"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	NumberPlain string    `db:"-" json:"number,omitempty"`
	ExpiryPlain string    `db:"-" json:"expiry,omitempty"`
}
