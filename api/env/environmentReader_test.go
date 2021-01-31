package env

import (
	"os"
	"testing"
)

func TestAllVariablesAreSet(t *testing.T) {
	_ = os.Setenv("SONNENBATTERIE_BASE_URL", "http://10.10.10.10:80")
	_ = os.Setenv("SONNENBATTERIE_USER_API_KEY", "test-user-auth-token")
	_ = os.Setenv("SONNENBATTERIE_SERVICE_API_KEY", "test-service-auth-token")

	tokens, err := ReadVariables()

	if err != nil {
		t.Errorf("ReadVariables() != %s, want nil", err.Error())
	}

	if tokens.BaseUrl != "http://10.10.10.10:80" {
		t.Errorf("ReadVariables().BaseUrl = %s; want http://10.10.10.10:80", tokens.BaseUrl)
	}

	if tokens.UserApiKey != "test-user-auth-token" {
		t.Errorf("ReadVariables().UserApiKey = %s; want test-user-auth-token", tokens.UserApiKey)
	}

	if tokens.ServiceApiKey != "test-service-auth-token" {
		t.Errorf("ReadVariables().ServiceApiKey = %s; want test-service-auth-token", tokens.ServiceApiKey)
	}
}

func TestServiceAuthTokenIsNotSet(t *testing.T) {
	_ = os.Setenv("SONNENBATTERIE_BASE_URL", "http://10.10.10.10:80")
	_ = os.Setenv("SONNENBATTERIE_USER_API_KEY", "test-user-auth-token")
	_ = os.Unsetenv("SONNENBATTERIE_SERVICE_API_KEY")

	tokens, err := ReadVariables()

	if err == nil {
		t.Errorf("ReadVariables() == nil, want error")
	}

	if tokens.BaseUrl != "" {
		t.Errorf("ReadVariables().BaseUrl = %s; want empty string", tokens.BaseUrl)
	}

	if tokens.UserApiKey != "" {
		t.Errorf("ReadVariables().UserApiKey = %s; want empty string", tokens.UserApiKey)
	}

	if tokens.ServiceApiKey != "" {
		t.Errorf("ReadVariables().ServiceApiKey = %s; want empty string", tokens.ServiceApiKey)
	}
}

func TestBaseUrlIsNotSet(t *testing.T) {
	_ = os.Unsetenv("SONNENBATTERIE_BASE_URL")
	_ = os.Setenv("SONNENBATTERIE_USER_API_KEY", "test-user-auth-token")
	_ = os.Setenv("SONNENBATTERIE_SERVICE_API_KEY", "test-service-auth-token")

	tokens, err := ReadVariables()

	if err == nil {
		t.Errorf("ReadVariables() == nil, want error")
	}

	if tokens.BaseUrl != "" {
		t.Errorf("ReadVariables().BaseUrl = %s; want empty string", tokens.BaseUrl)
	}

	if tokens.UserApiKey != "" {
		t.Errorf("ReadVariables().UserApiKey = %s; want empty string", tokens.UserApiKey)
	}

	if tokens.ServiceApiKey != "" {
		t.Errorf("ReadVariables().ServiceApiKey = %s; want empty string", tokens.ServiceApiKey)
	}
}

func TestUserAuthTokenIsNotSet(t *testing.T) {
	_ = os.Setenv("SONNENBATTERIE_BASE_URL", "http://10.10.10.10:80")
	_ = os.Unsetenv("SONNENBATTERIE_USER_API_KEY")
	_ = os.Setenv("SONNENBATTERIE_SERVICE_API_KEY", "test-service-auth-token")

	tokens, err := ReadVariables()

	if err == nil {
		t.Errorf("ReadVariables() == nil, want error")
	}

	if tokens.BaseUrl != "" {
		t.Errorf("ReadVariables().BaseUrl = %s; want empty string", tokens.BaseUrl)
	}

	if tokens.UserApiKey != "" {
		t.Errorf("ReadVariables().UserApiKey = %s; want empty string", tokens.UserApiKey)
	}

	if tokens.ServiceApiKey != "" {
		t.Errorf("ReadVariables().ServiceApiKey = %s; want empty string", tokens.ServiceApiKey)
	}
}

func TestTokensAreNotSet(t *testing.T) {
	_ = os.Unsetenv("SONNENBATTERIE_BASE_URL")
	_ = os.Unsetenv("SONNENBATTERIE_USER_API_KEY")
	_ = os.Unsetenv("SONNENBATTERIE_SERVICE_API_KEY")

	tokens, err := ReadVariables()

	if err == nil {
		t.Errorf("ReadVariables() == nil, want error")
	}

	if tokens.BaseUrl != "" {
		t.Errorf("ReadVariables().BaseUrl = %s; want empty string", tokens.BaseUrl)
	}

	if tokens.UserApiKey != "" {
		t.Errorf("ReadVariables().UserApiKey = %s; want empty string", tokens.UserApiKey)
	}

	if tokens.ServiceApiKey != "" {
		t.Errorf("ReadVariables().ServiceApiKey = %s; want empty string", tokens.ServiceApiKey)
	}
}
