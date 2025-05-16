package client

import (
	"sonnen-batterie-api/api/test"
	"strings"
	"testing"
)

func TestGetAuthToken(t *testing.T) {
	server := startSonnenBatterieServer(t)
	defer server.Close()
	ip := strings.Trim(server.URL, "http://")

	client := NewAuthClient(ip, "User", SonnenBatterieMockPassword)

	token, err := client.GetAuthToken()
	if err != nil {
		t.Fatalf("GetAuthToken() failed: %v", err)
	}

	test.Equals(t, SonnenBatterieMockAuthToken, token, "GetAuthToken()")
}

func TestGetAuthTokenError(t *testing.T) {
	// Test with invalid IP to trigger error
	client := NewAuthClient("invalid-ip", "User", SonnenBatterieMockPassword)

	_, err := client.GetAuthToken()
	if err == nil {
		t.Fatal("Expected error from GetAuthToken() with invalid IP, but got nil")
	}
}
