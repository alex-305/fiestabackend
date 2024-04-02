package main

import (
	"log"

	"github.com/alex-305/fiestabackend/models"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), err
}

func loginUser(creds models.Credentials, db *DB) error {
	dbPass, err := db.GetPassword(creds.Username)
	if err != nil {
		log.Printf("Error getting password: %v", err)
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(creds.Password))
	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}
	return nil
}
