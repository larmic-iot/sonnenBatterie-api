package client

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"sonnen-batterie-api/api/test"
	"testing"
)

const (
	apiKey = "cbe134b6-63c6-11eb-ae93-0242ac130002"
)

func TestClient_getLatestData(t *testing.T) {
	server := startServer(t)

	defer server.Close()

	client := NewClient(server.URL, apiKey)
	latestData, err := client.getLatestData()

	test.Equals(t, nil, err, "client.getLatestData()")
	test.Equals(t, 749, latestData.ConsumptionInWatt, "client.getLatestData()")
	test.Equals(t, 211, latestData.ProductionInWatt, "client.getLatestData()")
}

func TestClient_getLatestData_Api_Key_Not_Matching(t *testing.T) {
	server := startServer(t)

	defer server.Close()

	client := NewClient(server.URL, "not-matching")
	latestData, err := client.getLatestData()

	test.Equals(t, errors.New("status code is 401"), err, "client.getLatestData()")
	test.Equals(t, LatestData{}, latestData, "client.getLatestData()")
}

func TestClient_getLatestData_SonnenBatterie_not_found(t *testing.T) {
	client := NewClient("http://10.10.10.10:80", apiKey)
	latestData, err := client.getLatestData()

	timeoutError := mapToTimeoutError(err)

	test.Equals(t, "Get", timeoutError.Operation, "client.getLatestData()")
	test.Equals(t, "http://10.10.10.10:80/api/v2/latestdata", timeoutError.URL, "client.getLatestData()")
	test.Equals(t, "context deadline exceeded (Client.Timeout exceeded while awaiting headers)", timeoutError.Message, "client.getLatestData()")
	test.Equals(t, true, timeoutError.Timeout, "client.getLatestData()")
	test.Equals(t, LatestData{}, latestData, "client.getLatestData()")
}

func startServer(t *testing.T) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		test.Equals(t, req.URL.String(), "/api/v2/latestdata", "client.getLatestData()")

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
