package model

type EclipseLed struct {
	PulsingGreen  bool `json:"Pulsing Green"`
	PulsingOrange bool `json:"Pulsing Orange"`
	PulsingWhite  bool `json:"Pulsing White"`
	SolidRed      bool `json:"Solid Red"`
}

type IcStatus struct {
	EclipseLed EclipseLed `json:"Eclipse Led"`
}

type LatestDataDto struct {
	ConsumptionInWatt int64    `json:"Consumption_W"`
	ProductionInWatt  int64    `json:"Production_W"`
	IcStatus          IcStatus `json:"ic_status"`
}
