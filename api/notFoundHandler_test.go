package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestHandle404(t *testing.T) {
	response := httptest.NewRecorder()
	router := mux.NewRouter()
	router.NotFoundHandler = Handle404()

	router.ServeHTTP(response, httptest.NewRequest("GET", "/not-found", nil))

	if response.Code != http.StatusNotFound {
		t.Error("Did not get expected HTTP status code, got", response.Code)
	}

	if response.Header().Get("Content-Type") != "text/plain; charset=UTF-8" {
		t.Error("Did not get expected HTTP content type, got", response.Header().Get("Content-Type"))
	}

	if response.Body.String() != "{\"code\":404,\"message\":\"Not Found\"}\n" {
		t.Error("Did not get expected body, got", response.Body.String())
	}
}
