package model

import "encoding/json"

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

type SystemDto struct {
	BatterySystem BatterySystem `json:"battery_system"`
}
