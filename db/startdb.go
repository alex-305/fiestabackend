package db

import (
	"database/sql"
	"net/url"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func NewDB() (*DB, error) {
	/*Open the database*/
	dsn := url.URL{
		User:   url.UserPassword("postgres", "password"),
		Host:   "localhost:5432",
		Scheme: "postgres",
		Path:   "fiestapics",
	}
	db, err := sql.Open("postgres", dsn.String())
	if err != nil {
		return nil, err
	}

	/*Ping the database*/
	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
func (db *DB) Close() error {
	return db.DB.Close()
}
