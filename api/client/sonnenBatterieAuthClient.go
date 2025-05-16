package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sonnen-batterie-api/api/client/crypto"
	"strings"
)

type AuthClient struct {
	Ip       string
	User     string
	password string
}

type Token struct {
	Token string `json:"authentication_token"`
}

func NewAuthClient(ip, user, password string) *AuthClient {
	return &AuthClient{
		Ip:       ip,
		User:     user,
		password: password,
	}
}

func (c *AuthClient) GetAuthToken() (string, error) {
	baseUrl := "http://" + c.Ip + "/api/"

	challenge, err := GetChallenge(c.Ip)
	if err != nil {
		return "", fmt.Errorf("failed to get challenge: %w", err)
	}

	encryptedPassword := crypto.Encrypt(c.password, challenge)

	body := c.buildSessionBody(challenge, encryptedPassword)
	fmt.Println("POST session, Body: ", body)

	req, err := http.NewRequest("POST", baseUrl+"session", body)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	client := &http.Client{}
	post, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to execute request: %w", err)
	}
	defer post.Body.Close()

	token, err := unmarshalToken(post.Body)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal token: %w", err)
	}

	return token.Token, nil
}

func (c *AuthClient) buildSessionBody(challenge string, encryptedPassword string) io.Reader {
	values := url.Values{}
	values.Set("user", c.User)
	values.Set("challenge", challenge)
	values.Set("response", encryptedPassword)

	return strings.NewReader(values.Encode())
}

func (c *AuthClient) buildChallengeUrl() string {
	return "http://" + c.Ip + "/api/challenge"
}

func unmarshalToken(body io.ReadCloser) (Token, error) {
	token, err := io.ReadAll(body)
	if err != nil {
		return Token{}, fmt.Errorf("error while reading token: %w", err)
	}

	fmt.Println("Response from session:", string(token))

	var t Token
	err = json.Unmarshal(token, &t)
	if err != nil {
		return Token{}, fmt.Errorf("error while unmarshaling token: %w", err)
	}

	return t, nil
}
