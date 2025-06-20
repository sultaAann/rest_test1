package internal

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type ConfigDB struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

func ConnectDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://test:test@db:5432/test?sslmode=disable") // make norm config
	if err != nil {
		panic("Cannot connect DB")
	}

	err = db.Ping()
	if err != nil {
		panic("Cannot ping DB")
	}

	return db
}
