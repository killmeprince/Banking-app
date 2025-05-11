package database

import (
	"log"

	"banking-app/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitDB() *sqlx.DB {
	db, err := sqlx.Connect("postgres", config.GetDBConnStr())
	if err != nil {
		log.Fatalf("DB connect error: %v", err)
	}
	log.Println("Connected to Postgres")
	return db
}
