package api

import (
	"encoding/json"
	"net/http"
	sonnenBatterieClient "sonnen-batterie-api/api/client"
	model2 "sonnen-batterie-api/api/client/model"
	"sonnen-batterie-api/api/env"
	"sonnen-batterie-api/api/model"
)

func StatusHandler(env env.Environment, w http.ResponseWriter, _ *http.Request) {
	authClient := sonnenBatterieClient.NewAuthClient(env.Ip, env.UserName, env.UserPassword)
	token := authClient.GetAuthToken()

	client := sonnenBatterieClient.NewClient(env.Ip, token)
	statusDto, err := client.GetStatus()
	latestDataDto, err2 := client.GetLatestData()

	if err != nil || err2 != nil {
		w.Header().Add("Content-Type", "text/plain; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)

		_ = json.
			NewEncoder(w).
			Encode(
				model.ProtocolError{
					Code:    http.StatusNotFound,
					Message: env.Ip + " not found!",
				})
		return
	}

	system := model.Status{
		GridFeedIn:   statusDto.GridFeedInInWatt * 1000,
		ProductionMW: latestDataDto.ProductionInWatt * 1000,
		ChargeLevel:  int(statusDto.ChargeLevel),
		ChargeMode:   mapToChargeMode(statusDto),
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(system)
}

func mapToChargeMode(statusDto model2.StatusDto) model.ChargeMode {
	if statusDto.Charging {
		return model.Charging
	} else if statusDto.Discharging {
		return model.Discharging
	}

	return model.Idle
}
