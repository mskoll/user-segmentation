package repo

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	userTable        = "users"
	segmentTable     = "segment"
	userSegmentTable = "user_segment"
	operationTable   = "operation"
)

type Conf struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func Init(cfg Conf) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", connToString(cfg))
	// db, err := sqlx.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func connToString(info Conf) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		info.Host, info.Port, info.User, info.Password, info.DBName)
}
