package main

import (
	"database/sql"
	"log"
	"net/url"

	_ "github.com/lib/pq"

	"github.com/alex-305/fiestabackend/models"
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
func (db *DB) CreateUser(cred models.Credentials) error {
	var err error
	cred.Password, err = hashPassword(cred.Password)

	if err != nil {
		return err
	}

	stmt := `
	INSERT INTO users(username, password)
	VALUES($1,$2);`

	_, err = db.Exec(stmt, cred.Username, cred.Password)

	if err != nil {
		log.Printf("Error in insert: %v", err)
		return err
	}

	return nil
}

func (db *DB) GetPassword(username string) (string, error) {
	var password string
	err := db.QueryRow("SELECT password FROM users WHERE username = $1", username).Scan(&password)
	log.Printf("dbpass: %s", password)
	if err != nil {
		return "", err
	}

	return password, nil
}
