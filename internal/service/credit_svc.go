package service

import (
	"math"
	"time"

	"banking-app/internal/models"
	"banking-app/internal/repository"
)

type CreditService struct {
	repo   *repository.CreditRepo
	psRepo *repository.PSRepo
}

func NewCreditService(r *repository.CreditRepo, p *repository.PSRepo) *CreditService {
	return &CreditService{r, p}
}

func (s *CreditService) Create(accID int64, principal, annualRate float64, termMonths int) (*models.Credit, error) {
	monthlyRate := annualRate/12 + 0.05
	c := &models.Credit{
		AccountID:  accID,
		Principal:  principal,
		Rate:       monthlyRate,
		TermMonths: termMonths,
		Margin:     0.05,
	}
	if err := s.repo.Create(c); err != nil {
		return nil, err
	}

	schedule := computeAnnuitySchedule(c)
	if err := s.psRepo.CreateMany(schedule); err != nil {
		return nil, err
	}
	return c, nil
}

func (s *CreditService) GetSchedule(creditID int64) ([]models.PaymentSchedule, error) {
	return s.psRepo.ByCreditID(creditID)
}

func computeAnnuitySchedule(c *models.Credit) []models.PaymentSchedule {
	n := float64(c.TermMonths)
	r := c.Rate
	ann := c.Principal * (r * math.Pow(1+r, n)) / (math.Pow(1+r, n) - 1)
	var sched []models.PaymentSchedule
	for i := 1; i <= c.TermMonths; i++ {
		due := time.Now().AddDate(0, i, 0)
		sched = append(sched, models.PaymentSchedule{
			CreditID: c.ID,
			DueDate:  due,
			Amount:   math.Round(ann*100) / 100,
		})
	}
	return sched
}
