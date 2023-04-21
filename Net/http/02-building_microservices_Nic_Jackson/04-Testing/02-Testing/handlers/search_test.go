package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// validate that search criteria
// go test -v -race ./...
func TestSearchHandlerReturnsBadRequestWhenNoSearchCriteriaIsSent(t *testing.T) {
	//Arrange = Setup
	r, rw, handler := setupTest(nil)

	//Act = Execute
	handler.ServeHTTP(rw, r)

	//Assert
	if rw.Code != http.StatusBadRequest {
		t.Errorf("Expected BadRequest got %v", rw.Code)
	}
}

func TestSearchHandlerReturnsBadRequestWhenBlankSearchCriteriaIsSent(t *testing.T) {
	//Arrange = Setup
	r, rw, handler := setupTest(&searchRequest{})

	//Act = Execute
	handler.ServeHTTP(rw, r)

	//Assert
	if rw.Code != http.StatusBadRequest {
		t.Errorf("Expected BadRequest got %v", rw.Code)
	}
}

// Arrange = Setup
func setupTest(data interface{}) (*http.Request, *httptest.ResponseRecorder, Search) {
	/***** create HandlerSearch *****/
	handler := Search{}

	/***** create httptest.ResponseRecorder ~ http.ResponseWriter *****/
	res := httptest.NewRecorder()

	if data == nil {
		// httptest.NewRequest ~ http.Request
		return httptest.NewRequest("POST", "/search", nil), res, handler
	}

	/***** create httptest.NewRequest by passing some JSON in the request body *****/
	body, _ := json.Marshal(data)
	req := httptest.NewRequest("POST", "/search", bytes.NewReader(body))

	return req, res, handler
}
