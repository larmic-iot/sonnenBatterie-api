package model

type StatusDto struct {
	GridFeedInInWatt int64 `json:"GridFeedIn_W"`
	Charging         bool  `json:"BatteryCharging"`
	Discharging      bool  `json:"BatteryDischarging"`
	ChargeLevel      int64 `json:"RSOC"`
}
