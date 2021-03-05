package env

import (
	"errors"
	"os"
	"sonnen-batterie-api/api/test"
	"testing"
)

func TestAllVariablesAreSet(t *testing.T) {
	_ = os.Setenv("SONNENBATTERIE_IP", "10.10.10.10")
	_ = os.Setenv("SONNENBATTERIE_USER_NAME", "test-user-name")
	_ = os.Setenv("SONNENBATTERIE_USER_PASSWORD", "test-user-password")

	tokens, err := ReadVariables()

	test.Equals(t, nil, err, "ReadVariables()")
	test.Equals(t, tokens.Ip, "10.10.10.10", "ReadVariables().Ip")
	test.Equals(t, tokens.UserName, "test-user-name", "ReadVariables().UserName")
	test.Equals(t, tokens.UserPassword, "test-user-password", "ReadVariables().UserPassword")
}

func TestUserPasswordIsNotSet(t *testing.T) {
	_ = os.Setenv("SONNENBATTERIE_IP", "10.10.10.10")
	_ = os.Setenv("SONNENBATTERIE_USER_NAME", "test-user-name")
	_ = os.Unsetenv("SONNENBATTERIE_USER_PASSWORD")

	tokens, err := ReadVariables()

	test.Equals(t, errors.New("empty userName or userPassword"), err, "ReadVariables()")
	test.Equals(t, tokens.Ip, "", "ReadVariables().Ip")
	test.Equals(t, tokens.UserName, "", "ReadVariables().UserName")
	test.Equals(t, tokens.UserPassword, "", "ReadVariables().UserPassword")
}

func TestBaseUrlIsNotSet(t *testing.T) {
	_ = os.Unsetenv("SONNENBATTERIE_IP")
	_ = os.Setenv("SONNENBATTERIE_USER_NAME", "test-user-name")
	_ = os.Setenv("SONNENBATTERIE_USER_PASSWORD", "test-user-password")

	tokens, err := ReadVariables()

	test.Equals(t, errors.New("empty userName or userPassword"), err, "ReadVariables()")
	test.Equals(t, tokens.Ip, "", "ReadVariables().Ip")
	test.Equals(t, tokens.UserName, "", "ReadVariables().UserName")
	test.Equals(t, tokens.UserPassword, "", "ReadVariables().UserPassword")
}

func TestUserNameIsNotSet(t *testing.T) {
	_ = os.Setenv("SONNENBATTERIE_IP", "10.10.10.10")
	_ = os.Unsetenv("SONNENBATTERIE_USER_NAME")
	_ = os.Setenv("SONNENBATTERIE_USER_PASSWORD", "test-user-password")

	tokens, err := ReadVariables()

	test.Equals(t, errors.New("empty userName or userPassword"), err, "ReadVariables()")
	test.Equals(t, tokens.Ip, "", "ReadVariables().Ip")
	test.Equals(t, tokens.UserName, "", "ReadVariables().UserName")
	test.Equals(t, tokens.UserPassword, "", "ReadVariables().UserPassword")
}

func TestTokensAreNotSet(t *testing.T) {
	_ = os.Unsetenv("SONNENBATTERIE_IP")
	_ = os.Unsetenv("SONNENBATTERIE_USER_NAME")
	_ = os.Unsetenv("SONNENBATTERIE_USER_PASSWORD")

	tokens, err := ReadVariables()

	test.Equals(t, errors.New("empty userName or userPassword"), err, "ReadVariables()")
	test.Equals(t, tokens.Ip, "", "ReadVariables().Ip")
	test.Equals(t, tokens.UserName, "", "ReadVariables().UserName")
	test.Equals(t, tokens.UserPassword, "", "ReadVariables().UserPassword")
}
