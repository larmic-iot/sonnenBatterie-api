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

	test.Equals(t, int64(7801), data.ProductionInWatt, "GetLatestData()")
	test.Equals(t, int64(620), data.ConsumptionInWatt, "GetLatestData()")
}
