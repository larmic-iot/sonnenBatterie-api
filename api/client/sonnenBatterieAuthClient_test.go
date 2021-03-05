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

	client := NewAuthClient(ip, "User", SonnenBatterieMockPassword)

	test.Equals(t, SonnenBatterieMockAuthToken, client.GetAuthToken(), "GetAuthToken()")
}
