package routes

import (
	"encoding/json"
	"io"
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
	reader, writer := io.Pipe()

	body := NewGameBody{
		Name:  "testgame",
		Owner: "testuser",
	}

	go func() {
		defer writer.Close()

		if err := json.NewEncoder(writer).Encode(&body); err != nil {
			t.Errorf(err.Error())
		}
	}()

	req, err := http.NewRequest("POST", "", reader)

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(NewGameHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	response := Response{}
	if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
		t.Error(err)
	}
	if !response.Success {
		t.Errorf("handler returned unexpected body; success not true: got %v", response)
	}
}
