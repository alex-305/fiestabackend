package helpers

import (
	"errors"
	"net/http"
	"strings"
)

func GetToken(r *http.Request) (string, error) {

	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		return "", errors.New("missing authorization header")
	}

	parts := strings.Split(authHeader, " ")

	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("invalid authorization header format")
	}

	return parts[1], nil
}
