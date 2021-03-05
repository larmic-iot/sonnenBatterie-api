package env

import (
	"errors"
	"os"
)

type Environment struct {
	Ip           string
	UserName     string
	UserPassword string
}

func ReadVariables() (Environment, error) {
	var ip = os.Getenv("SONNENBATTERIE_IP")
	var userName = os.Getenv("SONNENBATTERIE_USER_NAME")
	var userPassword = os.Getenv("SONNENBATTERIE_USER_PASSWORD")

	if userName == "" || userPassword == "" || ip == "" {
		// TODO replace text
		return Environment{}, errors.New("empty userName or userPassword")
	}

	return Environment{
		Ip:           ip,
		UserName:     userName,
		UserPassword: userPassword,
	}, nil
}
