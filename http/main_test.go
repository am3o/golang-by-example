package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getProducts(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/api/v1/products", nil)
	response := httptest.NewRecorder()

	getProducts(response, request)

	var result []Product
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		t.Fail()
	}

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, products, result)
}
