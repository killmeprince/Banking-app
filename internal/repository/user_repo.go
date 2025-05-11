package repository

import (
	"banking-app/internal/models"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct{ db *sqlx.DB }

func NewUserRepo(db *sqlx.DB) *UserRepo { return &UserRepo{db} }

func (r *UserRepo) Create(u *models.User) error {
	_, err := r.db.NamedExec(
		`INSERT INTO users (email, password_hash) VALUES (:email, :password_hash)`, u)
	return err
}

func (r *UserRepo) ByEmail(email string) (*models.User, error) {
	u := &models.User{}
	err := r.db.Get(u,
		`SELECT id, email, password_hash, created_at FROM users WHERE email=$1`, email)
	return u, err
}
