package auth

import (
	"errors"
	"os"
	"time"

	"github.com/alex-305/fiestabackend/db"
	"github.com/alex-305/fiestabackend/models"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

type Claims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

func GenerateJWT(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("failure to create jwt claim")
	}
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	godotenv.Load()
	SecretKey := os.Getenv("SECRET_KEY")

	if SecretKey == "" {
		return "", errors.New("secret key could not be retrieved")
	}

	tokenString, err := token.SignedString([]byte(SecretKey))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(token string, db *db.DB) (models.User, error) {
	godotenv.Load()
	SecretKey := os.Getenv("SECRET_KEY")

	claims := &Claims{}

	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return models.User{}, err
	}

	user, err := db.GetUser(claims.Username)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
