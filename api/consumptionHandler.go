package api

import (
	"encoding/json"
	"net/http"
	sonnenBatterieClient "sonnen-batterie-api/api/client"
	"sonnen-batterie-api/api/env"
	"sonnen-batterie-api/api/model"
)

func ConsumptionHandler(env env.Environment, w http.ResponseWriter, _ *http.Request) {
	authClient := sonnenBatterieClient.NewAuthClient(env.Ip, env.UserName, env.UserPassword)
	token := authClient.GetAuthToken()

	client := sonnenBatterieClient.NewClient(env.Ip, token)
	latestData, err := client.GetLatestData()

	if err != nil {
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

	consumption := model.Consumption{
		CurrentMW: latestData.ConsumptionInWatt * 1000,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(consumption)
}
