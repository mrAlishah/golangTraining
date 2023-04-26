package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"

	"03-Testing/data"
)

var memStore *data.MemoryStore

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

func TestSearchHandlerCallsDataStoreWithValidQuery(t *testing.T) {
	//Arrange = Setup
	name := "Garfield"

	r, rw, handler := setupTest(&searchRequest{Query: name})

	//Act = Execute
	handler.ServeHTTP(rw, r)

	decoder := json.NewDecoder(rw.Body)
	response := new(searchResponse)
	err := decoder.Decode(response)

	//Assert
	if rw.Code == http.StatusBadRequest {
		t.Errorf("Expected Request Ok but got %v", rw.Code)
	}

	if err != nil {
		t.Errorf("Expected unMarshall response but got %v", err)
	}

	assert.Equal(t, 1, len(response.Kittens))
}

// Arrange = Setup
func setupTest(d interface{}) (*http.Request, *httptest.ResponseRecorder, Search) {
	memStore = &data.MemoryStore{}

	/***** create HandlerSearch *****/
	handler := Search{
		DataStore: memStore,
	}

	/***** create httptest.ResponseRecorder ~ http.ResponseWriter *****/
	res := httptest.NewRecorder()

	if d == nil {
		// httptest.NewRequest ~ http.Request
		return httptest.NewRequest("POST", "/search", nil), res, handler
	}

	/***** create httptest.NewRequest by passing some JSON in the request body *****/
	body, _ := json.Marshal(d)
	req := httptest.NewRequest("POST", "/search", bytes.NewReader(body))

	return req, res, handler
}
