package api

import (
	"net/http"
	"strconv"
	"strings"
)

func extractBearerToken(r *http.Request) int {

	authHeader := r.Header.Get("Authorization")

	tokens := strings.Split(authHeader, " ")

	if len(tokens) != 2 {
		return -1
	}

	token, err := strconv.Atoi(tokens[1])
	if err != nil {
		return -1
	}
	return token
}

func validToken(token int) bool {
	return token != -1
}

func validUsername(username string) bool {
	return len(username) >= 3 && len(username) <= 16
}

func validMessage(message string) bool {
	return len(message) >= 1 && len(message) <= 2200
}
