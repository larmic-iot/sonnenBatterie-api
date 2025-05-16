package client

import (
	"errors"
	"strings"
	"testing"
)

func TestGetChallenge(t *testing.T) {
	server := startSonnenBatterieServer(t)
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
		{"GetChallenge('127.0.0.1')", args{ip: "127.0.0.1"}, want{challenge: "", err: errors.New("could not acquire authentication challenge from sonnenBatterie: Get \"http://127.0.0.1/api/challenge\": dial tcp 127.0.0.1:80: connect: connection refused")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetChallenge(tt.args.ip)

			// Check challenge value
			if got != tt.want.challenge {
				t.Errorf("GetChallenge(%s) challenge = %v, want %v", tt.args.ip, got, tt.want.challenge)
			}

			// Check error
			if (err == nil && tt.want.err != nil) || (err != nil && tt.want.err == nil) {
				t.Errorf("GetChallenge(%s) error = %v, want %v", tt.args.ip, err, tt.want.err)
			} else if err != nil && tt.want.err != nil && err.Error() != tt.want.err.Error() {
				t.Errorf("GetChallenge(%s) error = %v, want %v", tt.args.ip, err.Error(), tt.want.err.Error())
			}
		})
	}
}
