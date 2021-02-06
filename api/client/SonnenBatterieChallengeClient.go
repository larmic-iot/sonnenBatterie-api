package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetChallenge(ip string) (string, error) {
	resp, err := http.Get(buildChallengeUrl(ip))

	if err != nil {
		fmt.Println("Could not acquire authentication challenge from sonnenBatterie", err)
		return "", err
	}

	defer resp.Body.Close()

	challengeBody, _ := ioutil.ReadAll(resp.Body)
	return trimChallenge(challengeBody), nil
}

func buildChallengeUrl(ip string) string {
	return "http://" + ip + "/api/challenge"
}

func trimChallenge(challenge []byte) string {
	return strings.Trim(string(challenge), "\"")
}
