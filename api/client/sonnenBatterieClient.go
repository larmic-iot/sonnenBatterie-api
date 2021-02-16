package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	contextPath = "/api/v2/latestdata"
)

type Client struct {
	Url    string
	apiKey string
}

type LatestData struct {
	ConsumptionInWatt int64 `json:"Consumption_W"`
	ProductionInWatt  int64 `json:"Production_W"`
}

func NewClient(baseUrl string, apiKey string) *Client {
	return &Client{
		Url:    baseUrl + contextPath,
		apiKey: apiKey,
	}
}

func (c *Client) GetLatestData() (LatestData, error) {
	client := &http.Client{Timeout: 1 * time.Second}

	req, err := http.NewRequest("GET", c.Url, nil)

	if err != nil {
		return LatestData{}, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Auth-Token", c.apiKey)
	resp, err := client.Do(req)

	if err != nil {
		return LatestData{}, err
	}

	if resp.StatusCode != 200 {
		return LatestData{}, errors.New(fmt.Sprintf("status code is %d", resp.StatusCode))
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return LatestData{}, err
	}

	var latestData LatestData

	err = json.Unmarshal(body, &latestData)

	if err != nil {
		return LatestData{}, err
	}

	return latestData, nil
}
