package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

const (
	plaintext = "text/plain; charset=utf-8"
)

func TestPingHandler_ShouldReturnAnswer(t *testing.T) {
	req, err := http.NewRequest("GET", "/ping", nil)
	assert.Nil(t, err, "Http request failed")
	recorder := httptest.NewRecorder()
	http.HandlerFunc(PingHandler).ServeHTTP(recorder, req)
	assert.Equal(t, plaintext, recorder.HeaderMap.Get("Content-type"))
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "pong", string(recorder.Body.Bytes()))
}

func TestLoginHandler_WithValidBody_ShouldReturnToken(t *testing.T) {
	login := LoginData{Email: "demo@empatica.com", Password: "passw0rd"}
	loginjson, _ := json.Marshal(login)
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(loginjson))
	assert.Nil(t, err, "Http request failed")
	recorder := httptest.NewRecorder()
	http.HandlerFunc(LoginHandler).ServeHTTP(recorder, req)
	assert.Equal(t, plaintext, recorder.HeaderMap.Get("Content-type"))
	assert.Equal(t, http.StatusOK, recorder.Code)
	var ua *UserAccount
	err2 := json.Unmarshal(recorder.Body.Bytes(), &ua)
	assert.Nil(t, err2)
	assert.Equal(t, "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9", ua.Token)
}

func TestLoginHandler_WithInvalidUser_Should403Forbidden(t *testing.T) {
	login := LoginData{Email: "invaliduser", Password: "passw0rd"}
	loginjson, _ := json.Marshal(login)
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(loginjson))
	assert.Nil(t, err, "Http request failed")
	recorder := httptest.NewRecorder()
	http.HandlerFunc(LoginHandler).ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusForbidden, recorder.Code)
	assert.Equal(t, "invalid user", string(recorder.Body.Bytes()))
}

func TestLoginHandler_WithInvalidBody_ShouldReturn400BadRequest(t *testing.T) {
	req, err := http.NewRequest("POST", "/login", strings.NewReader("invalidjson"))
	assert.Nil(t, err, "Http request failed")
	recorder := httptest.NewRecorder()
	http.HandlerFunc(LoginHandler).ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Equal(t, "invalid login data", string(recorder.Body.Bytes()))
}

func TestGetUserHandler_WithExistingId_ShouldReturnUserDetails(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/{userId}", nil)
	req = mux.SetURLVars(req, map[string]string{"userId": "1"})
	assert.Nil(t, err, "Http request failed")
	recorder := httptest.NewRecorder()
	http.HandlerFunc(GetUserHandler).ServeHTTP(recorder, req)
	assert.Equal(t, plaintext, recorder.HeaderMap.Get("Content-type"))
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "{\"id\":1,\"email\":\"demo@empatica.com\",\"firstName\":\"John\",\"lastName\":\"\",\"age\":13}\n", string(recorder.Body.Bytes()))
}

func TestGetUserHandler_WithUnexistingId_ShouldReturn404NotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/{userId}", nil)
	req = mux.SetURLVars(req, map[string]string{"userId": "47"})
	assert.Nil(t, err, "Http request failed")
	recorder := httptest.NewRecorder()
	http.HandlerFunc(GetUserHandler).ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusNotFound, recorder.Code)
	assert.Equal(t, "missing user", string(recorder.Body.Bytes()))
}

func TestGetUserHandler_WithMalformedUrl_ShouldReturn400BadRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/{userId}", nil)
	req = mux.SetURLVars(req, map[string]string{"userId": "invalidint"})
	assert.Nil(t, err, "Http request failed")
	recorder := httptest.NewRecorder()
	http.HandlerFunc(GetUserHandler).ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}
