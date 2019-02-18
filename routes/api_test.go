package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(HealthHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"status":"healthy"}`
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestNewGameHandler(t *testing.T) {

}
