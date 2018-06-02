package main

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUser_WithValidId_ShouldReturnUser(t *testing.T) {
	var (
		userID int64 = 1
		user   *User
	)
	user = getUser(userID)
	assert.NotNil(t, user)
	assert.Equal(t, user.ID, userID)
	assert.Equal(t, user.Email, "demo@empatica.com")
	assert.Equal(t, user.FirstName, "John")
	assert.Equal(t, user.LastName, "") //user is underage
	assert.Equal(t, user.Age, 13)
}

func TestGetUser_WithInvalidId_ShouldReturnNil(t *testing.T) {
	var (
		userID int64 = 23
		user   *User
	)
	user = getUser(userID)
	assert.Nil(t, user)
}

func IsBase64(s string) bool {
	_, err := base64.StdEncoding.DecodeString(s)
	return err == nil
}

func TestGetUserAccount_WithValidId_ShouldReturnUserAccount(t *testing.T) {
	var (
		userID int64 = 1
		ua     *UserAccount
	)
	ua = getUserAccount(userID)
	assert.NotNil(t, ua)
	assert.Equal(t, ua.UserID, userID)
	assert.Equal(t, ua.Password, "passw0rd")
	assert.True(t, IsBase64(ua.Token))
}

func TestGetUserAccount_WithInvalidId_ShouldReturnNil(t *testing.T) {
	var (
		userID int64 = -12
		ua     *UserAccount
	)
	ua = getUserAccount(userID)
	assert.Nil(t, ua)
}

func TestGetLogin_WithValidCredentials_ShouldReturnUserAccount(t *testing.T) {

	var (
		username = "demo@empatica.com"
		password = "passw0rd"
		ua       *UserAccount
		err      error
	)
	ua, err = getLogin(username, password)
	assert.NotNil(t, ua)
	assert.Nil(t, err)
}

func TestGetLogin_WithNotFoundUser_ShouldReturnError(t *testing.T) {

	var (
		username = "invaliduser"
		password = "passw0rd"
		ua       *UserAccount
		err      error
	)
	ua, err = getLogin(username, password)
	assert.Nil(t, ua)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "invalid user")
}

func TestGetLogin_WithNotMatchingPassword_ShouldReturnError(t *testing.T) {

	var (
		username = "demo@empatica.com"
		password = "invalidpassword"
		ua       *UserAccount
		err      error
	)
	ua, err = getLogin(username, password)
	assert.Nil(t, ua)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "invalid password")
}

var sanitizeTests = []struct {
	user            *User
	expectedsurname string
}{
	{
		&User{ID: 135, Email: "someEmail", FirstName: "firstName", LastName: "lastName", Age: 19},
		"lastName",
	},
	{
		&User{ID: 135, Email: "someEmail", FirstName: "firstName", LastName: "lastName", Age: 7},
		"",
	},
}

func TestSanitize(t *testing.T) {

	for _, tt := range sanitizeTests {
		var newuser = tt.user
		newuser.Sanitize()
		assert.Equal(t, newuser.LastName, tt.expectedsurname)
	}

}
