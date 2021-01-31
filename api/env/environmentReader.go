package env

import (
	"errors"
	"os"
)

type Environment struct {
	User    string
	Service string
}

func ReadVariables() (Environment, error) {
	var userAuthToken = os.Getenv("USER_AUTH_TOKEN")
	var serviceAuthToken = os.Getenv("SERVICE_AUTH_TOKEN")

	if userAuthToken == "" || serviceAuthToken == "" {
		return Environment{}, errors.New("empty userAuthToken or serviceAuthToken")
	}

	return Environment{
		User:    userAuthToken,
		Service: serviceAuthToken,
	}, nil
}
