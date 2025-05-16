package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"sonnen-batterie-api/api/client/crypto"
	"testing"
)

type body struct {
	User      string `json:"user"`
	Challenge string `json:"challenge"`
	Response  string `json:"response"`
}

const (
	SonnenBatterieMockChallenge = "test-challenge"
	SonnenBatterieMockPassword  = "test-password"
	SonnenBatterieMockAuthToken = "test-auth-token"
)

func startSonnenBatterieServer(t *testing.T) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		switch req.URL.String() {
		case "/api/challenge":
			doGetChallenge(t, rw, req)
			break

		case "/api/session":
			doPostSession(t, rw, req)

			break

		case "/api/v2/latestdata":
			doGetLatestData(t, rw, req)

			break

		case "/api/v2/status":
			doGetStatus(t, rw, req)

			break

		case "/api/battery_system":
			doGetSystem(t, rw, req)

			break
		default:
			t.Fatalf("URL %s not known on test server", req.URL.String())
		}
	}))

	return server
}

func doGetChallenge(t *testing.T, rw http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		_ = fmt.Sprintf("%s /api/session, want: GET", req.Method)
		t.FailNow()
	}

	_, _ = rw.Write([]byte("\"" + SonnenBatterieMockChallenge + "\""))
}

func doPostSession(t *testing.T, rw http.ResponseWriter, req *http.Request) {
	encryptedPassword := crypto.Encrypt(SonnenBatterieMockPassword, SonnenBatterieMockChallenge)

	if req.Method != "POST" {
		t.Fatalf("%s /api/session, want: POST", req.Method)
	}

	err := req.ParseForm()
	if err != nil {
		t.Fatalf("Failed to parse form: %v", err)
	}

	user := req.Form.Get("user")
	challenge := req.Form.Get("challenge")
	response := req.Form.Get("response")

	if user != "User" {
		t.Fatalf("POST /api/session form user: = %s, want: User", user)
	}
	if challenge != SonnenBatterieMockChallenge {
		t.Fatalf("POST /api/session form challenge: = %s, want: %s", challenge, SonnenBatterieMockChallenge)
	}
	if response != encryptedPassword {
		t.Fatalf("POST /api/session form response: = %s, want: %s", response, encryptedPassword)
	}

	rw.Header().Set("Content-Type", "application/json")
	_, err = rw.Write([]byte("{\"authentication_token\":\"" + SonnenBatterieMockAuthToken + "\"}"))
	if err != nil {
		t.Fatalf("Failed to write response: %v", err)
	}
}

func doGetLatestData(t *testing.T, rw http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		t.Fatalf("%s /api/v2/latestdata, want: GET", req.Method)
	}

	dat, err := os.ReadFile("test_response_get_latest_data.json")
	if err != nil {
		t.Fatalf("Failed to read test_response_get_latest_data.json: %v", err)
	}

	rw.WriteHeader(http.StatusOK)
	_, err = rw.Write(dat)
	if err != nil {
		t.Fatalf("Failed to write response: %v", err)
	}
}

func doGetStatus(t *testing.T, rw http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		t.Fatalf("%s /api/v2/status, want: GET", req.Method)
	}

	dat, err := os.ReadFile("test_response_get_status.json")
	if err != nil {
		t.Fatalf("Failed to read test_response_get_status.json: %v", err)
	}

	rw.WriteHeader(http.StatusOK)
	_, err = rw.Write(dat)
	if err != nil {
		t.Fatalf("Failed to write response: %v", err)
	}
}

func doGetSystem(t *testing.T, rw http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		t.Fatalf("%s /api/battery_system, want: GET", req.Method)
	}

	dat, err := os.ReadFile("test_response_get_battery_system.json")
	if err != nil {
		t.Fatalf("Failed to read test_response_get_battery_system.json: %v", err)
	}

	rw.WriteHeader(http.StatusOK)
	_, err = rw.Write(dat)
	if err != nil {
		t.Fatalf("Failed to write response: %v", err)
	}
}
