package api

import (
	"errors"
	"os"
)

type AuthenticationToken struct {
	User    string
	Service string
}

func ReadAuthenticationTokens() (AuthenticationToken, error) {
	var userAuthToken = os.Getenv("USER_AUTH_TOKEN")
	var serviceAuthToken = os.Getenv("SERVICE_AUTH_TOKEN")

	if userAuthToken == "" || serviceAuthToken == "" {
		return AuthenticationToken{}, errors.New("empty userAuthToken or serviceAuthToken")
	}

	return AuthenticationToken{
		User:    userAuthToken,
		Service: serviceAuthToken,
	}, nil
}
