package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {

	//Initializtion

	//Execution
	user, err := GetUser(0)

	//Validation
	assert.Nil(t, user, "we are not expecting a user with id 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not found", err.Code)
	assert.EqualValues(t, "user 0 does not exists", err.Message)

}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Nathan", user.FirstName)
	assert.EqualValues(t, "Keeks", user.LastName)
	assert.EqualValues(t, "test123@test.com", user.Email)
}
