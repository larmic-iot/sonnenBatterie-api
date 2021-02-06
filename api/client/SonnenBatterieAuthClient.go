package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sonnen-batterie-api/api/client/crypto"
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

	// get token
	body := `{"user":"` + c.User + `","challenge":"` + challenge + `","response":"` + encryptedPassword + `"}`
	fmt.Println("POST session, Body: ", body)
	post, _ := http.Post(baseUrl+"session", "application/json", bytes.NewBuffer([]byte(body)))
	defer post.Body.Close()
	token, _ := ioutil.ReadAll(post.Body)
	var t Token
	_ = json.Unmarshal(token, &t)
	return t.Token
}

func getSecure(url, token string) string {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Auth-Token", token)
	client := &http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func (c *AuthClient) buildChallengeUrl() string {
	return "http://" + c.Ip + "/api/challenge"
}
