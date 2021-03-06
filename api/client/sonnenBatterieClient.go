package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Ip    string
	token string
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

type Status struct {
	GridFeedInInWatt int64 `json:"GridFeedIn_W"`
	Charging         bool  `json:"BatteryCharging"`
	Discharging      bool  `json:"BatteryDischarging"`
	ChargeLevel      int64 `json:"RSOC"`
}

type Network struct {
	WanIP string `json:"wanip"`
	LanIP string `json:"lanip"`
	VpnIP string `json:"vpnip"`
}

type Software struct {
	FirmwareVersion string `json:"firmware_version"`
}

type BatterySystemSystem struct {
	MacAddress      string      `json:"mac_address"`
	ModelName       string      `json:"model_name"`
	HardwareVersion json.Number `json:"hardware_version"`
}

type BatterySystem struct {
	Network             Network             `json:"network"`
	Software            Software            `json:"software"`
	BatterySystemSystem BatterySystemSystem `json:"system"`
}

type System struct {
	BatterySystem BatterySystem `json:"battery_system"`
}

func NewClient(ip string, user string, password string) *Client {
	client := NewAuthClient(ip, user, password)
	token := client.GetAuthToken()

	return &Client{
		Ip:    ip,
		token: token,
	}
}

func (c *Client) GetLatestData() (LatestData, error) {
	var latestData LatestData

	err := c.getRequest("/api/v2/latestdata", &latestData)

	return latestData, err
}

func (c *Client) GetStatus() (Status, error) {
	var status Status

	err := c.getRequest("/api/v2/status", &status)

	return status, err
}

func (c *Client) GetSystem() (System, error) {
	var system System

	err := c.getRequest("/api/battery_system", &system)

	return system, err
}

func (c *Client) getRequest(urlContextPath string, dto interface{}) error {
	url := "http://" + c.Ip + urlContextPath
	response := getRequest(url, c.token)

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
