package model

type Led string

const (
	White  Led = "white"
	Green      = "green"
	Orange     = "orange"
	Red        = "red"
)

type System struct {
	Ip              string `json:"ip"`
	WanIp           string `json:"wanIp"`
	Model           string `json:"model"`
	MacAddress      string `json:"mac"`
	SoftwareVersion string `json:"softwareVersion"`
	HardwareVersion string `json:"hardwareVersion"`
	Led             Led    `json:"led"`
}
