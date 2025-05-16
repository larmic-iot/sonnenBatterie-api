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

func (c *AuthClient) GetAuthToken() string {
	baseUrl := "http://" + c.Ip + "/api/"

	challenge, _ := GetChallenge(c.Ip)
	encryptedPassword := crypto.Encrypt(c.password, challenge)

	body := c.buildSessionBody(challenge, encryptedPassword)
	fmt.Println("POST session, Body: ", body)
	req, _ := http.NewRequest("POST", baseUrl+"session", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	client := &http.Client{}
	post, _ := client.Do(req)
	defer post.Body.Close()

	return unmarshalToken(post.Body).Token
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

func unmarshalToken(body io.ReadCloser) Token {
	token, err := io.ReadAll(body)
	if err != nil {
		fmt.Println("Error while reading token:", err)
		return Token{}
	}

	fmt.Println("Response from challenge:", string(token))

	var t Token
	_ = json.Unmarshal(token, &t)
	return t
}
