package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sonnen-batterie-api/api/client/crypto"
	"sonnen-batterie-api/api/test"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	password := "test-password"

	server := startSonnenBatterieServer(t, password)
	defer server.Close()
	ip := strings.Trim(server.URL, "http://")

	client := NewAuthClient(ip, "User", password)

	test.Equals(t, "test-auth-token", client.GetAuthToken(), "GetAuthToken()")
}

func startSonnenBatterieServer(t *testing.T, password string) *httptest.Server {
	challenge := "test-challenge"
	encryptedPassword := crypto.Encrypt(password, challenge)

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" && req.URL.String() == "/api/challenge" {
			_, _ = rw.Write([]byte("\"" + challenge + "\""))
		} else if req.Method == "POST" && req.URL.String() == "/api/session" {
			var b body
			_ = json.NewDecoder(req.Body).Decode(&b)

			if b.User != "User" {
				_ = fmt.Sprintf("POST /api/session body user: = %s, want: User", b.User)
				t.FailNow()
			}
			if b.Challenge != challenge {
				_ = fmt.Sprintf("POST /api/session body challenge: = %s, want: "+challenge+"", b.Challenge)
				t.FailNow()
			}
			if b.Response != encryptedPassword {
				_ = fmt.Sprintf("POST /api/session body response: = %s, want: "+encryptedPassword+"", b.Response)
				t.FailNow()
			}

			// TODO verify body
			_, _ = rw.Write([]byte("{\"authentication_token\":\"test-auth-token\"}"))
		} else {
			_ = fmt.Sprintf("URL %s not known on test server", req.URL.String())
			t.FailNow()
		}
	}))

	return server
}
