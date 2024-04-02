package main

import (
	"database/sql"
	"net/url"

	"github.com/alex-305/fiestabackend/models"
)

type DB struct {
	*sql.DB
}

func NewDB() (*DB, error) {
	/*Open the database*/
	dsn := url.URL{
		User:   url.UserPassword("username", "password"),
		Host:   "localhost:1433",
		Scheme: "postgres",
		Path:   "dbname",
	}
	db, err := sql.Open("postgres", dsn.String())
	if err != nil {
		return nil, err
	}

	defer db.Close()

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
	VALUES($1,$2)`

	_, err = db.Exec(stmt, cred.Username, cred.Password)

	if err != nil {
		return err
	}

	return nil
}

func (s *APIServer) getUser(username, password string) (*models.Credentials, error) {

	var user models.Credentials
	var err error
	//err := db.QueryRow("SELECT id, username, password FROM users WHERE username = $1", username).Scan(&user.ID, &user.username, &user.password)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return &user, err
}
