package api

import (
	"encoding/json"
	"net/http"
	sonnenBatterieClient "sonnen-batterie-api/api/client"
	model2 "sonnen-batterie-api/api/client/model"
	"sonnen-batterie-api/api/env"
	"sonnen-batterie-api/api/model"
)

func SystemHandler(env env.Environment, w http.ResponseWriter, _ *http.Request) {
	client := sonnenBatterieClient.NewClient(env.Ip, env.UserName, env.UserPassword)
	systemDto, err := client.GetSystem()
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

	system := model.System{
		Ip:              systemDto.BatterySystem.Network.LanIP,
		WanIp:           systemDto.BatterySystem.Network.WanIP,
		Model:           systemDto.BatterySystem.BatterySystemSystem.ModelName,
		MacAddress:      systemDto.BatterySystem.BatterySystemSystem.MacAddress,
		SoftwareVersion: systemDto.BatterySystem.Software.FirmwareVersion,
		HardwareVersion: string(systemDto.BatterySystem.BatterySystemSystem.HardwareVersion),
		Led:             mapToLed(latestDataDto.IcStatus.EclipseLed),
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(system)
}

func mapToLed(eclipseLed model2.EclipseLed) model.Led {
	if eclipseLed.SolidRed {
		return model.Red
	} else if eclipseLed.PulsingOrange {
		return model.Orange
	} else if eclipseLed.PulsingGreen {
		return model.Green
	}

	return model.White
}
