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

func (c *Client) GetStatus() (Status, error) {
	client := NewAuthClient(c.Ip, c.user, c.password)
	token := client.GetAuthToken()

	url := "http://" + c.Ip + "/api/v2/status"

	response := getRequest(url, token)

	var status Status

	err := json.Unmarshal([]byte(response), &status)

	if err != nil {
		return Status{}, err
	}

	return status, nil
}

func (c *Client) GetSystem() (System, error) {
	client := NewAuthClient(c.Ip, c.user, c.password)
	token := client.GetAuthToken()

	url := "http://" + c.Ip + "/api/battery_system"

	response := getRequest(url, token)

	var system System

	err := json.Unmarshal([]byte(response), &system)

	if err != nil {
		return System{}, err
	}

	return system, nil
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
