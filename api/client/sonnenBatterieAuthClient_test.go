package client

import (
	"sonnen-batterie-api/api/test"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	server := startSonnenBatterieServer(t)
	defer server.Close()
	ip := strings.Trim(server.URL, "http://")

	client := NewAuthClient(ip, "User", Password)

	test.Equals(t, "test-auth-token", client.GetAuthToken(), "GetAuthToken()")
}
