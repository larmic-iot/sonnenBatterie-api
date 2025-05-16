package client

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func GetChallenge(ip string) (string, error) {
	resp, err := http.Get(buildChallengeUrl(ip))
	if err != nil {
		return "", fmt.Errorf("could not acquire authentication challenge from sonnenBatterie: %w", err)
	}

	defer resp.Body.Close()

	challengeBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading challenge response: %w", err)
	}

	return trimChallenge(challengeBody), nil
}

func buildChallengeUrl(ip string) string {
	return "http://" + ip + "/api/challenge"
}

func trimChallenge(challenge []byte) string {
	return strings.Trim(string(challenge), "\"")
}
