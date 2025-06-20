package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewDB(cfg Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	))
	if err != nil {
		errMessage := fmt.Sprintf("Error opening DB connetion: %v", err)
		return nil, errors.New(errMessage)
	}

	if err = db.Ping(); err != nil {
		errMessage := fmt.Sprintf("Ping() DB error: %v", err)
		return nil, errors.New(errMessage)
	}
	return db, nil
}
