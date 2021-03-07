package client

import (
	"encoding/json"
	"sonnen-batterie-api/api/test"
	"strings"
	"testing"
)

//goland:noinspection GoNilness
func TestGetLatestData(t *testing.T) {
	server := startSonnenBatterieServer(t)
	defer server.Close()
	ip := strings.Trim(server.URL, "http://")

	client := NewClient(ip, SonnenBatterieMockAuthToken)
	data, err := client.GetLatestData()

	if err != nil {
		t.Errorf("GetLatestData(%s) returns error", ip)
	}

	test.Equals(t, int64(7135), data.ProductionInWatt, "GetLatestData().ProductionInWatt")
	test.Equals(t, int64(675), data.ConsumptionInWatt, "GetLatestData().ConsumptionInWatt")
	test.Equals(t, false, data.IcStatus.EclipseLed.PulsingGreen, "GetLatestData().IcStatus.EclipseLed.PulsingGreen")
	test.Equals(t, false, data.IcStatus.EclipseLed.PulsingOrange, "GetLatestData().IcStatus.EclipseLed.PulsingOrange")
	test.Equals(t, true, data.IcStatus.EclipseLed.PulsingWhite, "GetLatestData().IcStatus.EclipseLed.PulsingWhite")
	test.Equals(t, false, data.IcStatus.EclipseLed.SolidRed, "GetLatestData().IcStatus.EclipseLed.SolidRed")
}

//goland:noinspection GoNilness
func TestGetStatus(t *testing.T) {
	server := startSonnenBatterieServer(t)
	defer server.Close()
	ip := strings.Trim(server.URL, "http://")

	client := NewClient(ip, SonnenBatterieMockAuthToken)
	data, err := client.GetStatus()

	if err != nil {
		t.Errorf("GetStatus(%s) returns error", ip)
	}

	test.Equals(t, int64(82), data.ChargeLevel, "GetStatus().ChargeLevel")
	test.Equals(t, int64(-35), data.GridFeedInInWatt, "GetStatus().GridFeedInInWatt")
	test.Equals(t, false, data.Charging, "GetStatus().Charging")
	test.Equals(t, true, data.Discharging, "GetStatus().Discharging")
}

//goland:noinspection GoNilness
func TestGetSystem(t *testing.T) {
	server := startSonnenBatterieServer(t)
	defer server.Close()
	ip := strings.Trim(server.URL, "http://")

	client := NewClient(ip, SonnenBatterieMockAuthToken)
	data, err := client.GetSystem()

	if err != nil {
		t.Errorf("GetSystem(%s) returns error", ip)
	}

	test.Equals(t, "10.0.0.100", data.BatterySystem.Network.LanIP, "GetSystem().BatterySystem.Network.LanIP")
	test.Equals(t, "31.31.31.31", data.BatterySystem.Network.WanIP, "GetSystem().BatterySystem.Network.WanIP")
	test.Equals(t, "41.41.41.41", data.BatterySystem.Network.VpnIP, "GetSystem().BatterySystem.Network.VpnIP")
	test.Equals(t, "1.5.5.781986", data.BatterySystem.Software.FirmwareVersion, "GetSystem().BatterySystem.Software.FirmwareVersion")
	test.Equals(t, "power unit sB10s sI1 9010 IP30", data.BatterySystem.BatterySystemSystem.ModelName, "GetSystem().BatterySystem.BatterySystemSystem.ModelName")
	test.Equals(t, json.Number("10.0"), data.BatterySystem.BatterySystemSystem.HardwareVersion, "GetSystem().BatterySystem.BatterySystemSystem.HardwareVersion")
}
