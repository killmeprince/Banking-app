package service

import (
	"errors"
	"fmt"
	"time"

	"banking-app/config"
	"banking-app/internal/models"
	"banking-app/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{ repo *repository.UserRepo }

func NewUserService(r *repository.UserRepo) *UserService { return &UserService{r} }

func (s *UserService) Register(email, pass string) error {
	if _, err := s.repo.ByEmail(email); err == nil {
		return errors.New("user already exists")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u := &models.User{Email: email, PasswordHash: string(h)}
	return s.repo.Create(u)
}

func (s *UserService) Login(email, pass string) (string, error) {
	u, err := s.repo.ByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}
	if bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(pass)) != nil {
		return "", errors.New("invalid credentials")
	}
	claims := jwt.RegisteredClaims{
		Subject:   fmt.Sprint(u.ID),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tok.SignedString(config.JWTSecret)
}
