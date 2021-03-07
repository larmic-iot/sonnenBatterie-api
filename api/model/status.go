package model

type ChargeMode string

const (
	Charging    ChargeMode = "charging"
	Discharging            = "discharging"
	Idle                   = "idle"
)

type Status struct {
	GridFeedIn   int64      `json:"grid_feed_in_mw"`
	ProductionMW int64      `json:"production_mw"`
	ChargeLevel  int        `json:"charge_level"`
	ChargeMode   ChargeMode `json:"charge_mode"`
}
