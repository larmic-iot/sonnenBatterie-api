package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sonnen-batterie-api/api/client/crypto"
)

const (
	JsonContentType = "application/json"
)

type AuthClient struct {
	Ip       string
	User     string
	password string
}

type Token struct {
	Token string `json:"authentication_token"`
}

type body struct {
	User      string `json:"User"`
	Challenge string `json:"challenge"`
	Response  string `json:"response"`
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
	post, _ := http.Post(baseUrl+"session", JsonContentType, body)
	defer post.Body.Close()

	return unmarshalToken(post.Body).Token
}

func (c *AuthClient) buildSessionBody(challenge string, encryptedPassword string) io.Reader {
	body := `{"user":"` + c.User + `","challenge":"` + challenge + `","response":"` + encryptedPassword + `"}`
	return bytes.NewBuffer([]byte(body))
}

func (c *AuthClient) buildChallengeUrl() string {
	return "http://" + c.Ip + "/api/challenge"
}

func unmarshalToken(body io.ReadCloser) Token {
	token, _ := ioutil.ReadAll(body)
	var t Token
	_ = json.Unmarshal(token, &t)
	return t
}
