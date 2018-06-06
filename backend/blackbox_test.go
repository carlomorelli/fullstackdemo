//+build blackbox

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	baseurl string
)

// Expect and env variable called BASEURL. If not defined, use the default "localhost:9000";
// The Env variable is handy for multiple deployments and CI server
func TestMain(m *testing.M) {
	baseurl = os.Getenv("BASEURL")
	if baseurl == "" {
		baseurl = "localhost:9000"
	}
	fmt.Printf("Testing on baseUrl=%s\n", baseurl)
	m.Run()
}

func TestPing(t *testing.T) {
	res, err := http.Get("http://" + baseurl + "/ping")
	assert.Nil(t, err, "Error in HTTP connection")
	status, body := extract(t, res)
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, "pong", body)
}

func TestCorsNotAllowedMethod(t *testing.T) {
	fmt.Printf("Test cors with not allowed method")
	req, _ := http.NewRequest("PATCH", "http://"+baseurl+"/ping", nil)
	res, err := http.DefaultClient.Do(req)
	assert.Nil(t, err, "Error in HTTP connection")
	status, body := extract(t, res)
	assert.Equal(t, http.StatusMethodNotAllowed, status)
	assert.Equal(t, "", body)

}

// In general a blackbox test should not have knowledge of the internal data structure,
// Here I'm making an exception by using LoginData struct from the code and marshalling to json
var loginTests = []struct {
	testcasedescription string
	logindata           *LoginData
	expectedstatus      int
	expectedbody        string
}{
	{
		"Happy path, test a working login credentials, expect token in answer\n",
		&LoginData{Email: "demo@empatica.com", Password: "passw0rd"},
		http.StatusOK,
		"{\"token\":\"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9\"}\n",
	},
	{
		"Not valid login credentials, expect error\n",
		&LoginData{Email: "demo@empatica.com", Password: "wrongpassword"},
		http.StatusForbidden,
		"invalid password",
	},
}

func TestLogin(t *testing.T) {
	for _, tt := range loginTests {
		fmt.Printf(tt.testcasedescription)
		jsonlogin, _ := json.Marshal(tt.logindata)
		buf := bytes.NewBuffer([]byte(jsonlogin))
		res, err := http.Post("http://"+baseurl+"/login", "application/json", buf)
		assert.Nil(t, err, "Error in HTTP connection")
		status, body := extract(t, res)
		assert.Equal(t, tt.expectedstatus, status)
		assert.Equal(t, tt.expectedbody, body)
	}
}

func TestLoginWithMalformedBody(t *testing.T) {
	fmt.Printf("Malformed login credentials, expect error\n")
	buf := bytes.NewBuffer([]byte("not a valid json"))
	res, err := http.Post("http://"+baseurl+"/login", "application/json", buf)
	assert.Nil(t, err, "Error in HTTP connection")
	status, body := extract(t, res)
	assert.Equal(t, http.StatusBadRequest, status)
	assert.Equal(t, "invalid login data", body)

}

func TestGetAllUsers(t *testing.T) {
	// functionality not implemented, so expecting error
	res, err := http.Get("http://" + baseurl + "/users")
	assert.Nil(t, err, "Error in HTTP connection")
	status, body := extract(t, res)
	assert.Equal(t, http.StatusNotFound, status)
	assert.Equal(t, "404 page not found\n", body)

}

func TestGetUnderageUser(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://"+baseurl+"/users/1", nil)
	req.Header.Add("Authorization", "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9")
	res, err := http.DefaultClient.Do(req)
	assert.Nil(t, err, "Error in HTTP connection")
	status, body := extract(t, res)
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, "{\"id\":1,\"email\":\"demo@empatica.com\",\"firstName\":\"John\",\"lastName\":\"\",\"age\":13}\n", body)
	//lastname is correctly empty as user is underage
}

func TestGetUserWithoutToken(t *testing.T) {
	res, err := http.Get("http://" + baseurl + "/users/1")
	assert.Nil(t, err, "Error in HTTP connection")
	status, body := extract(t, res)
	assert.Equal(t, http.StatusUnauthorized, status)
	assert.Equal(t, "\"missing token\"\n", body)
}

func TestGetUnexistingUser(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://"+baseurl+"/users/718", nil)
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9")
	res, err := http.DefaultClient.Do(req)
	assert.Nil(t, err, "Error in HTTP connection")
	status, body := extract(t, res)
	assert.Equal(t, http.StatusForbidden, status)
	assert.Equal(t, "\"invalid token\"\n", body)
}

// Utility
func extract(t *testing.T, res *http.Response) (int, string) {
	defer res.Body.Close()
	body, ioerr := ioutil.ReadAll(res.Body)
	assert.Nil(t, ioerr, "Malformed body received in HTTP connection")
	return res.StatusCode, string(body)
}
