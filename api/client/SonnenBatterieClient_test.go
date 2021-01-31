package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestClient_getLatestData(t *testing.T) {
	dat, err := ioutil.ReadFile("test_response.json")

	fmt.Print(dat)

	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		equals(t, req.URL.String(), "/api/v2/latestdata")
		// Send response to be tested
		rw.Write(dat)
	}))

	// Close the server when test finishes
	defer server.Close()

	client := NewClient("http://10.0.40.100:80", "cbe134b6-63c6-11eb-ae93-0242ac130002")
	latestData, err := client.getLatestData()

	fmt.Println(latestData.ConsumptionInWatt)
	fmt.Println(latestData.ProductionInWatt)
	fmt.Println(err)
}

func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
