package client

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"sonnen-batterie-api/api/test"
	"strings"
	"testing"
)

func TestGetChallenge(t *testing.T) {
	server := startChallengeServer(t)
	defer server.Close()
	ip := strings.Trim(server.URL, "http://")

	type args struct{ ip string }
	type want struct {
		challenge string
		err       error
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{"GetChallenge('" + ip + "')", args{ip: ip}, want{challenge: "test-challenge", err: nil}},
		{"GetChallenge('127.0.0.1')", args{ip: "127.0.0.1"}, want{challenge: "", err: errors.New("Get \"http://127.0.0.1/api/challenge\": dial tcp 127.0.0.1:80: connect: connection refused")}},
		// TODO test does not work on github actions
		//{"GetChallenge('no-ip')", args{ip: "no-ip"}, want{challenge: "", err: errors.New("Get \"http://no-ip/api/challenge\": dial tcp: lookup no-ip: no such host")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := GetChallenge(tt.args.ip); got != tt.want.challenge {
				t.Errorf("GetChallenge("+tt.args.ip+") = %v, want %v", got, tt.want.challenge)
			}
			if _, err := GetChallenge(tt.args.ip); err != nil && err.Error() != tt.want.err.Error() {
				t.Errorf("GetChallenge("+tt.args.ip+") = %v, want %v", err.Error(), tt.want.err.Error())
			}
		})
	}
}

func startChallengeServer(t *testing.T) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		test.Equals(t, req.URL.String(), "/api/challenge", "GetChallenge(ip)")

		_, _ = rw.Write([]byte("\"test-challenge\""))
	}))

	return server
}
