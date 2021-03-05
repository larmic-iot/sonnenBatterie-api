package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Ip       string
	user     string
	password string
}

type EclipseLed struct {
	PulsingGreen  bool `json:"Pulsing Green"`
	PulsingOrange bool `json:"Pulsing Orange"`
	PulsingWhite  bool `json:"Pulsing White"`
	SolidRed      bool `json:"Solid Red"`
}

type IcStatus struct {
	EclipseLed EclipseLed `json:"Eclipse Led"`
}

type LatestData struct {
	ConsumptionInWatt int64    `json:"Consumption_W"`
	ProductionInWatt  int64    `json:"Production_W"`
	IcStatus          IcStatus `json:"ic_status"`
}

func NewClient(ip string, user string, password string) *Client {
	return &Client{
		Ip:       ip,
		user:     user,
		password: password,
	}
}

func (c *Client) GetLatestData() (LatestData, error) {
	client := NewAuthClient(c.Ip, c.user, c.password)
	token := client.GetAuthToken()

	url := "http://" + c.Ip + "/api/v2/latestdata"

	response := getRequest(url, token)

	var latestData LatestData

	err := json.Unmarshal([]byte(response), &latestData)

	if err != nil {
		return LatestData{}, err
	}

	return latestData, nil
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
