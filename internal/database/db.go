package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"userSegmentation/config"
)

type DB struct {
	DB *sqlx.DB
}

func New(cfg *config.DB) (*DB, error) {

	db, err := sqlx.Open("postgres", dbUrl(*cfg))

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func dbUrl(cfg config.DB) string {

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
}
