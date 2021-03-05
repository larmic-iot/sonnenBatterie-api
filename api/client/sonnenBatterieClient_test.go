package client

import (
	"sonnen-batterie-api/api/test"
	"strings"
	"testing"
)

//goland:noinspection GoNilness
func TestGetLatestData(t *testing.T) {
	server := startSonnenBatterieServer(t)
	defer server.Close()
	ip := strings.Trim(server.URL, "http://")

	client := NewClient(ip, "User", SonnenBatterieMockPassword)
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
