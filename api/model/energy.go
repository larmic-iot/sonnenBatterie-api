package model

type Consumption struct {
	CurrentMW int64 `json:"current_mw"`
	TodayWH   int64 `json:"today_wh"`
	TotalWH   int64 `json:"total_wh"`
}
