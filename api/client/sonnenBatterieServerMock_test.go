package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sonnen-batterie-api/api/client/crypto"
	"testing"
)

const (
	Challenge = "test-challenge"
	Password  = "test-password"
	AuthToken = "test-auth-token"
)

func startSonnenBatterieServer(t *testing.T) *httptest.Server {
	encryptedPassword := crypto.Encrypt(Password, Challenge)

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" && req.URL.String() == "/api/challenge" {
			_, _ = rw.Write([]byte("\"" + Challenge + "\""))
		} else if req.Method == "POST" && req.URL.String() == "/api/session" {
			var b body
			_ = json.NewDecoder(req.Body).Decode(&b)

			if b.User != "User" {
				_ = fmt.Sprintf("POST /api/session body user: = %s, want: User", b.User)
				t.FailNow()
			}
			if b.Challenge != Challenge {
				_ = fmt.Sprintf("POST /api/session body challenge: = %s, want: "+Challenge+"", b.Challenge)
				t.FailNow()
			}
			if b.Response != encryptedPassword {
				_ = fmt.Sprintf("POST /api/session body response: = %s, want: "+encryptedPassword+"", b.Response)
				t.FailNow()
			}

			_, _ = rw.Write([]byte("{\"authentication_token\":\"" + AuthToken + "\"}"))
		} else {
			_ = fmt.Sprintf("URL %s not known on test server", req.URL.String())
			t.FailNow()
		}
	}))

	return server
}
