package tests

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"mta-hosting-optimizer/handlers"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetHostnames(t *testing.T) {
	os.Setenv("THRESHOLD", "1")

	router := gin.Default()
	router.GET("/hostnames", handlers.GetHostnames)

	req, err := http.NewRequest("GET", "/hostnames", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := []string{"mta-prod-1", "mta-prod-3"}
	var result []string
	if err := json.NewDecoder(recorder.Body).Decode(&result); err != nil {
		t.Fatal(err)
	}

	if !compareSlices(result, expected) {
		t.Errorf("Handler returned unexpected body: got %v want %v", result, expected)
	}
}

func compareSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	return true
}
