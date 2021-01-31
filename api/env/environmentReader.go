package env

import (
	"errors"
	"os"
)

type Environment struct {
	BaseUrl       string
	UserApiKey    string
	ServiceApiKey string
}

func ReadVariables() (Environment, error) {
	var baseUrl = os.Getenv("SONNENBATTERIE_BASE_URL")
	var userApiKey = os.Getenv("SONNENBATTERIE_USER_API_KEY")
	var serviceApiKey = os.Getenv("SONNENBATTERIE_SERVICE_API_KEY")

	if userApiKey == "" || serviceApiKey == "" || baseUrl == "" {
		// TODO replace text
		return Environment{}, errors.New("empty userApiKey or serviceApiKey")
	}

	return Environment{
		BaseUrl:       baseUrl,
		UserApiKey:    userApiKey,
		ServiceApiKey: serviceApiKey,
	}, nil
}
