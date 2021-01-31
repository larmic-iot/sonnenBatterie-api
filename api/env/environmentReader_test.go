package env

import (
	"errors"
	"os"
	"sonnen-batterie-api/api/test"
	"testing"
)

func TestAllVariablesAreSet(t *testing.T) {
	_ = os.Setenv("SONNENBATTERIE_BASE_URL", "http://10.10.10.10:80")
	_ = os.Setenv("SONNENBATTERIE_USER_API_KEY", "test-user-auth-token")
	_ = os.Setenv("SONNENBATTERIE_SERVICE_API_KEY", "test-service-auth-token")

	tokens, err := ReadVariables()

	test.Equals(t, nil, err, "ReadVariables()")
	test.Equals(t, tokens.BaseUrl, "http://10.10.10.10:80", "ReadVariables().BaseUrl")
	test.Equals(t, tokens.UserApiKey, "test-user-auth-token", "ReadVariables().UserApiKey")
	test.Equals(t, tokens.ServiceApiKey, "test-service-auth-token", "ReadVariables().ServiceApiKey")
}

func TestServiceAuthTokenIsNotSet(t *testing.T) {
	_ = os.Setenv("SONNENBATTERIE_BASE_URL", "http://10.10.10.10:80")
	_ = os.Setenv("SONNENBATTERIE_USER_API_KEY", "test-user-auth-token")
	_ = os.Unsetenv("SONNENBATTERIE_SERVICE_API_KEY")

	tokens, err := ReadVariables()

	test.Equals(t, errors.New("empty userApiKey or serviceApiKey"), err, "ReadVariables()")
	test.Equals(t, tokens.BaseUrl, "", "ReadVariables().BaseUrl")
	test.Equals(t, tokens.UserApiKey, "", "ReadVariables().UserApiKey")
	test.Equals(t, tokens.ServiceApiKey, "", "ReadVariables().ServiceApiKey")
}

func TestBaseUrlIsNotSet(t *testing.T) {
	_ = os.Unsetenv("SONNENBATTERIE_BASE_URL")
	_ = os.Setenv("SONNENBATTERIE_USER_API_KEY", "test-user-auth-token")
	_ = os.Setenv("SONNENBATTERIE_SERVICE_API_KEY", "test-service-auth-token")

	tokens, err := ReadVariables()

	test.Equals(t, errors.New("empty userApiKey or serviceApiKey"), err, "ReadVariables()")
	test.Equals(t, tokens.BaseUrl, "", "ReadVariables().BaseUrl")
	test.Equals(t, tokens.UserApiKey, "", "ReadVariables().UserApiKey")
	test.Equals(t, tokens.ServiceApiKey, "", "ReadVariables().ServiceApiKey")
}

func TestUserAuthTokenIsNotSet(t *testing.T) {
	_ = os.Setenv("SONNENBATTERIE_BASE_URL", "http://10.10.10.10:80")
	_ = os.Unsetenv("SONNENBATTERIE_USER_API_KEY")
	_ = os.Setenv("SONNENBATTERIE_SERVICE_API_KEY", "test-service-auth-token")

	tokens, err := ReadVariables()

	test.Equals(t, errors.New("empty userApiKey or serviceApiKey"), err, "ReadVariables()")
	test.Equals(t, tokens.BaseUrl, "", "ReadVariables().BaseUrl")
	test.Equals(t, tokens.UserApiKey, "", "ReadVariables().UserApiKey")
	test.Equals(t, tokens.ServiceApiKey, "", "ReadVariables().ServiceApiKey")
}

func TestTokensAreNotSet(t *testing.T) {
	_ = os.Unsetenv("SONNENBATTERIE_BASE_URL")
	_ = os.Unsetenv("SONNENBATTERIE_USER_API_KEY")
	_ = os.Unsetenv("SONNENBATTERIE_SERVICE_API_KEY")

	tokens, err := ReadVariables()

	test.Equals(t, errors.New("empty userApiKey or serviceApiKey"), err, "ReadVariables()")
	test.Equals(t, tokens.BaseUrl, "", "ReadVariables().BaseUrl")
	test.Equals(t, tokens.UserApiKey, "", "ReadVariables().UserApiKey")
	test.Equals(t, tokens.ServiceApiKey, "", "ReadVariables().ServiceApiKey")
}
