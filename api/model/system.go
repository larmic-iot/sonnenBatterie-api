package model

type Switch string

const (
	On  Switch = "on"
	Off        = "off"
)

type System struct {
	Ip              string `json:"ip"`
	Port            int    `json:"port"`
	Name            string `json:"name"`
	Icon            string `json:"icon"`
	Model           string `json:"model"`
	MacAddress      string `json:"mac"`
	SoftwareVersion string `json:"softwareVersion"`
	HardwareVersion string `json:"hardwareVersion"`
	Led             Switch `json:"led"`
	Power           Switch `json:"power"`
}
