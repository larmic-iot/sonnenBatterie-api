package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"sonnen-batterie-api/api/client/crypto"
	"testing"
)

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
			_ = fmt.Sprintf("URL %s not known on test server", req.URL.String())
			t.FailNow()
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
		_ = fmt.Sprintf("%s /api/session, want: POST", req.Method)
		t.FailNow()
	}

	var b body
	_ = json.NewDecoder(req.Body).Decode(&b)

	if b.User != "User" {
		_ = fmt.Sprintf("POST /api/session body user: = %s, want: User", b.User)
		t.FailNow()
	}
	if b.Challenge != SonnenBatterieMockChallenge {
		_ = fmt.Sprintf("POST /api/session body challenge: = %s, want: "+SonnenBatterieMockChallenge+"", b.Challenge)
		t.FailNow()
	}
	if b.Response != encryptedPassword {
		_ = fmt.Sprintf("POST /api/session body response: = %s, want: "+encryptedPassword+"", b.Response)
		t.FailNow()
	}

	_, _ = rw.Write([]byte("{\"authentication_token\":\"" + SonnenBatterieMockAuthToken + "\"}"))
}

func doGetLatestData(t *testing.T, rw http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		_ = fmt.Sprintf("%s /api/v2/latestdata, want: GET", req.Method)
		t.FailNow()
	}

	dat, _ := ioutil.ReadFile("test_response_get_latest_data.json")
	rw.WriteHeader(http.StatusOK)
	_, _ = rw.Write(dat)
}

func doGetStatus(t *testing.T, rw http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		_ = fmt.Sprintf("%s /api/v2/status, want: GET", req.Method)
		t.FailNow()
	}

	dat, _ := ioutil.ReadFile("test_response_get_status.json")
	rw.WriteHeader(http.StatusOK)
	_, _ = rw.Write(dat)
}

func doGetSystem(t *testing.T, rw http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		_ = fmt.Sprintf("%s /api/battery_system, want: GET", req.Method)
		t.FailNow()
	}

	dat, _ := ioutil.ReadFile("test_response_get_battery_system.json")
	rw.WriteHeader(http.StatusOK)
	_, _ = rw.Write(dat)
}
