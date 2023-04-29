package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"

	"04-Testing/data"
)

var mockStore *data.MockStore

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

	// look at repo:https://github.com/stretchr/testify#mock-package
	/*
	  // create an instance of our test object
	  testObj := new(MyMockedObject)

	  // setup expectations
	  testObj.On("DoSomething", 123).Return(true, nil)
	*/
	// Result made fake by .Return(make([]data.Kitten, 0))
	mockStore.On("Search", name).Return(make([]data.Kitten, 0))

	//Act = Execute
	handler.ServeHTTP(rw, r)

	//Assert
	mockStore.AssertExpectations(t)
}

func TestSearchHandlerReturnsKittensWithValidQuery(t *testing.T) {
	//Arrange = Setup
	name := "Fat Freddy's Cat"
	r, rw, handler := setupTest(&searchRequest{Query: name})

	// Result made fake by .Return(make([]data.Kitten, 1))
	mockStore.On("Search", name).Return(make([]data.Kitten, 1))

	//Act = Execute
	handler.ServeHTTP(rw, r)

	response := searchResponse{}
	json.Unmarshal(rw.Body.Bytes(), &response)

	//Assert
	// we have fake Mock result => .Return(make([]data.Kitten, 1))
	t.Log(response.Kittens)

	assert.Equal(t, 1, len(response.Kittens))
	assert.Equal(t, http.StatusOK, rw.Code)
}

// Arrange = Setup
func setupTest(d interface{}) (*http.Request, *httptest.ResponseRecorder, Search) {
	//memStore = &data.MemoryStore{}
	mockStore = &data.MockStore{}

	/***** create HandlerSearch *****/
	handler := Search{
		//DataStore: memStore,
		DataStore: mockStore,
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
