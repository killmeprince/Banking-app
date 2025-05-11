package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	"banking-app/config"
	"banking-app/internal/models"
	"banking-app/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

func luhnGenerate() string {
	digits := make([]int, 15)
	for i := range digits {
		digits[i] = rand.Intn(10)
	}
	sum := 0
	for i, d := range digits {
		v := d
		if i%2 == 0 {
			v *= 2
			if v > 9 {
				v -= 9
			}
		}
		sum += v
	}
	check := (10 - sum%10) % 10
	s := ""
	for _, d := range digits {
		s += fmt.Sprint(d)
	}
	return s + fmt.Sprint(check)
}

type CardService struct {
	repo *repository.CardRepo
}

func NewCardService(r *repository.CardRepo) *CardService { return &CardService{r} }

func (s *CardService) Create(accountID int64) (*models.Card, error) {
	num := luhnGenerate()
	exp := time.Now().AddDate(3, 0, 0).Format("2006-01-02")
	c := &models.Card{
		AccountID:   accountID,
		NumberPlain: num,
		ExpiryPlain: exp,
		NumberEnc:   []byte(num),
		ExpEnc:      []byte(exp),
	}

	cvv := fmt.Sprintf("%03d", rand.Intn(1000))
	hash, _ := bcrypt.GenerateFromPassword([]byte(cvv), bcrypt.DefaultCost)
	hm := hmac.New(sha256.New, config.JWTSecret)
	hm.Write([]byte(num))
	mac := hex.EncodeToString(hm.Sum(nil))

	c.CVVHash = string(hash)
	c.HMAC = mac

	return c, s.repo.Create(c)
}
