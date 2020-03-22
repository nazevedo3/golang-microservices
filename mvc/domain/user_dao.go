package domain

import (
	"fmt"
	"net/http"

	"github.com/nazevedo3/golang-microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Nathan", LastName: "Keeks", Email: "test123@test.com"},
	}
)

func GetUser(userID int64) (*User, *utils.ApplicationError) {
	if user := users[userID]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v does not exists", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not found",
	}

}
