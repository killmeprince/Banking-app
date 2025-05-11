package models

import "time"

type Account struct {
	ID        int64     `db:"id" json:"id"`
	UserID    int64     `db:"user_id" json:"user_id"`
	Balance   float64   `db:"balance" json:"balance"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
