package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	plaintext = "text/plain; charset=utf-8"
)

func TestPingHandler_ShouldReturnAnswer(t *testing.T) {

	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		assert.Fail(t, "Http request failed")
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(PingHandler)
	handler.ServeHTTP(recorder, req)

	assert.Equal(t, plaintext, recorder.HeaderMap.Get("Content-type"))
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "pong", string(recorder.Body.Bytes()))

}

func TestGetUserHandler_WithValidLoginInput_ShouldReturnValidUserDetails(t *testing.T) {

	req, err := http.NewRequest("GET", "/users/1", nil)
	if err != nil {
		assert.Fail(t, "Http request failed")
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUserHandler)
	handler.ServeHTTP(recorder, req)

	assert.Equal(t, plaintext, recorder.HeaderMap.Get("Content-type"))
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "pong", string(recorder.Body.Bytes()))

}

/* //mock function
func getUser(userID int64) *User {

	args := m.Called(userID)
	return args.User

}
*/
