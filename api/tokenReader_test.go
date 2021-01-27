package api

import (
	"os"
	"testing"
)

func TestBothTokensAreSet(t *testing.T) {
	_ = os.Setenv("USER_AUTH_TOKEN", "test-user-auth-token")
	_ = os.Setenv("SERVICE_AUTH_TOKEN", "test-service-auth-token")

	tokens, err := ReadAuthenticationTokens()

	if err != nil {
		t.Errorf("ReadAuthenticationTokens() != %s, want nil", err.Error())
	}

	if tokens.User != "test-user-auth-token" {
		t.Errorf("ReadAuthenticationTokens().User = %s; want test-user-auth-token", tokens.User)
	}

	if tokens.Service != "test-service-auth-token" {
		t.Errorf("ReadAuthenticationTokens().Service = %s; want test-service-auth-token", tokens.Service)
	}
}

func TestServiceAuthTokenIsNotSet(t *testing.T) {
	_ = os.Setenv("USER_AUTH_TOKEN", "test-user-auth-token")
	_ = os.Unsetenv("SERVICE_AUTH_TOKEN")

	tokens, err := ReadAuthenticationTokens()

	if err == nil {
		t.Errorf("ReadAuthenticationTokens() == nil, want error")
	}

	if tokens.User != "" {
		t.Errorf("ReadAuthenticationTokens().User = %s; want empty string", tokens.User)
	}

	if tokens.Service != "" {
		t.Errorf("ReadAuthenticationTokens().Service = %s; want empty string", tokens.Service)
	}
}

func TestUserAuthTokenIsNotSet(t *testing.T) {
	_ = os.Unsetenv("USER_AUTH_TOKEN")
	_ = os.Setenv("SERVICE_AUTH_TOKEN", "test-service-auth-token")

	tokens, err := ReadAuthenticationTokens()

	if err == nil {
		t.Errorf("ReadAuthenticationTokens() == nil, want error")
	}

	if tokens.User != "" {
		t.Errorf("ReadAuthenticationTokens().User = %s; want empty string", tokens.User)
	}

	if tokens.Service != "" {
		t.Errorf("ReadAuthenticationTokens().Service = %s; want empty string", tokens.Service)
	}
}

func TestTokensAreNotSet(t *testing.T) {
	_ = os.Unsetenv("USER_AUTH_TOKEN")
	_ = os.Unsetenv("SERVICE_AUTH_TOKEN")

	tokens, err := ReadAuthenticationTokens()

	if err == nil {
		t.Errorf("ReadAuthenticationTokens() == nil, want error")
	}

	if tokens.User != "" {
		t.Errorf("ReadAuthenticationTokens().User = %s; want empty string", tokens.User)
	}

	if tokens.Service != "" {
		t.Errorf("ReadAuthenticationTokens().Service = %s; want empty string", tokens.Service)
	}
}
