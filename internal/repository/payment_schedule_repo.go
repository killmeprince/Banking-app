package repository

import (
	"banking-app/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
)

type PSRepo struct{ db *sqlx.DB }

func NewPSRepo(db *sqlx.DB) *PSRepo { return &PSRepo{db} }

func (r *PSRepo) CreateMany(schedules []models.PaymentSchedule) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	for _, s := range schedules {
		if _, err := tx.NamedExec(
			`INSERT INTO payment_schedules (credit_id, due_date, amount)
			 VALUES (:credit_id, :due_date, :amount)`,
			s,
		); err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func (r *PSRepo) ByCreditID(creditID int64) ([]models.PaymentSchedule, error) {
	var sched []models.PaymentSchedule
	err := r.db.Select(
		&sched,
		`SELECT credit_id, due_date, amount 
		 FROM payment_schedules 
		 WHERE credit_id=$1 
		 ORDER BY due_date`,
		creditID,
	)
	return sched, err
}

func (r *PSRepo) ByCredit(crID int64) ([]models.PaymentSchedule, error) {
	var xs []models.PaymentSchedule
	err := r.db.Select(&xs,
		`SELECT * FROM payment_schedules WHERE credit_id=$1 ORDER BY due_date`, crID)
	return xs, err
}

func (r *PSRepo) MarkPaid(id int64) error {
	_, err := r.db.Exec(
		`UPDATE payment_schedules SET paid=true WHERE id=$1`, id)
	return err
}

func (r *PSRepo) UnpaidDue(now time.Time) ([]struct {
	ID        int64
	AccountID int64
	Amount    float64
}, error) {
	var xs []struct {
		ID        int64
		AccountID int64
		Amount    float64
	}
	err := r.db.Select(&xs, `
		SELECT ps.id, c.account_id, ps.amount
		FROM payment_schedules ps
		JOIN credits c ON ps.credit_id = c.id
		WHERE ps.paid=false AND ps.due_date<= $1
	`, now)
	return xs, err
}
