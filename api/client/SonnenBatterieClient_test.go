package client

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

const (
	apiKey = "cbe134b6-63c6-11eb-ae93-0242ac130002"
)

func TestClient_getLatestData(t *testing.T) {
	server := startServer(t)

	defer server.Close()

	client := NewClient(server.URL, apiKey, 0)
	latestData, err := client.getLatestData()

	equals(t, nil, err)
	equals(t, 749, latestData.ConsumptionInWatt)
	equals(t, 211, latestData.ProductionInWatt)
}

func TestClient_getLatestData_Api_Key_Not_Matching(t *testing.T) {
	server := startServer(t)

	defer server.Close()

	client := NewClient(server.URL, "not-matching", 0)
	latestData, err := client.getLatestData()

	equals(t, errors.New("status code is 401"), err)
	equals(t, LatestData{}, latestData)
}

func TestClient_getLatestData_SonnenBatterie_not_found(t *testing.T) {
	client := NewClient("http://10.10.10.10:80", apiKey, 10)
	latestData, err := client.getLatestData()

	timeoutError := mapToTimeoutError(err)

	equals(t, "Get", timeoutError.Operation)
	equals(t, "http://10.10.10.10:80/api/v2/latestdata", timeoutError.URL)
	equals(t, "context deadline exceeded (Client.Timeout exceeded while awaiting headers)", timeoutError.Message)
	equals(t, true, timeoutError.Timeout)
	equals(t, LatestData{}, latestData)
}

func startServer(t *testing.T) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		equals(t, req.URL.String(), "/api/v2/latestdata")

		if req.Header.Get("Auth-Token") != apiKey {
			rw.WriteHeader(http.StatusUnauthorized)
			_, _ = rw.Write([]byte("Status code is 401"))
		} else {
			dat, _ := ioutil.ReadFile("test_response.json")
			rw.WriteHeader(http.StatusOK)
			_, _ = rw.Write(dat)
		}
	}))

	return server
}

func mapToTimeoutError(err error) *TimeoutError {
	urlError := err.(*url.Error)
	val := reflect.ValueOf(urlError.Err).Elem()

	return &TimeoutError{
		Operation: urlError.Op,
		URL:       urlError.URL,
		Message:   val.Field(0).String(),
		Timeout:   val.Field(1).Bool(),
	}
}

type TimeoutError struct {
	Operation string
	URL       string
	Message   string
	Timeout   bool
}

func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
