package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearchHandlerReturnsBadRequestWhenNoSearchCriteriaIsSent(t *testing.T) {
	handler := Search{}
	request := httptest.NewRequest("GET", "/search", nil)
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected BadRequest got %v", response.Code)
	}
}
