package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/nazevedo3/golang-microservices/mvc/utils"

	"github.com/nazevedo3/golang-microservices/mvc/services"
)

func GetUser(w http.ResponseWriter, r *http.Request) {

	userId, err := strconv.ParseInt(r.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		jsonValue, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		return

	}

	user, apiErr := services.GetUser(userId)

	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write([]byte(jsonValue))

		// Handle the err and return to the client
		return
	}

	// return user to client

	jsonValue, _ := json.Marshal(user)
	w.Write(jsonValue)

}
