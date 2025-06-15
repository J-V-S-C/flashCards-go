package repository

import (
	"database/sql"
	"errors"
	"fmt"
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
	db, err := sql.Open("pkg", fmt.Sprintf("user=%v password=%v host=%v port=%v database=%v sslmode=disable", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName))
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
