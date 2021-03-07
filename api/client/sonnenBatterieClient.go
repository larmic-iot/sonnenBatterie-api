package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sonnen-batterie-api/api/client/model"
)

type Client struct {
	Ip    string
	token string
}

func NewClient(ip string, user string, password string) *Client {
	client := NewAuthClient(ip, user, password)
	token := client.GetAuthToken()

	return &Client{
		Ip:    ip,
		token: token,
	}
}

func (c *Client) GetLatestData() (model.LatestDataDto, error) {
	var latestData model.LatestDataDto

	err := c.getRequest("/api/v2/latestdata", &latestData)

	return latestData, err
}

func (c *Client) GetStatus() (model.StatusDto, error) {
	var status model.StatusDto

	err := c.getRequest("/api/v2/status", &status)

	return status, err
}

func (c *Client) GetSystem() (model.SystemDto, error) {
	var system model.SystemDto

	err := c.getRequest("/api/battery_system", &system)

	return system, err
}

func (c *Client) getRequest(urlContextPath string, dto interface{}) error {
	url := "http://" + c.Ip + urlContextPath
	response := getRequest(url, c.token)

	fmt.Println("token " + c.token)
	fmt.Printf("GET %s, Body: %s\n", url, response)

	return json.Unmarshal([]byte(response), &dto)
}

func getRequest(url, token string) string {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Auth-Token", token)
	client := &http.Client{}
	// TODO return err
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
