package config

import (
	"fmt"
	"os"
	"strconv"
)

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

var (
	DBHost = getEnv("DB_HOST", "localhost")
	DBPort = getEnv("DB_PORT", "5432")
	DBUser = getEnv("DB_USER", "user")
	DBPass = getEnv("DB_PASS", "pass")
	DBName = getEnv("DB_NAME", "banking")

	JWTSecret = []byte(getEnv("JWT_SECRET", "supersecretkey"))

	SMTPHost = getEnv("SMTP_HOST", "smtp.example.com")
	SMTPPort = func() int {
		p, _ := strconv.Atoi(getEnv("SMTP_PORT", "587"))
		return p
	}()
	SMTPUser = getEnv("SMTP_USER", "noreply@example.com")
	SMTPPass = getEnv("SMTP_PASS", "strong_password")
)

func GetDBConnStr() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		DBUser, DBPass, DBHost, DBPort, DBName,
	)
}
