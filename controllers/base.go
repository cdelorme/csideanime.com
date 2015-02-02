package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Base struct {
	HandleBadRequest          func(writer http.ResponseWriter, request *http.Request)
	HandleInternalServerError func(writer http.ResponseWriter, request *http.Request)
}

// bad request handling
func (base *Base) BadRequest(writer http.ResponseWriter, request *http.Request) {
	if base.HandleBadRequest != nil {
		base.HandleBadRequest(writer, request)
		return
	}
	message := struct{ Error string }{Error: "Bad Request: " + request.URL.String() + ", using request method " + request.Method}
	jsonMessage, _ := json.Marshal(message)
	writer.WriteHeader(http.StatusBadRequest)
	writer.Header().Add("Content-Type", "application/json")
	writer.Write(jsonMessage)
}

func (base *Base) StartOperations() {
	// return uncacheable objects used in all operations
	// authentication for member login/checkin, or guest model
	// access control list with groups & actions
	//
}

// internal server error handling
func (base *Base) InternalServerError(writer http.ResponseWriter, request *http.Request) {
	if base.HandleInternalServerError != nil {
		base.HandleInternalServerError(writer, request)
		return
	}
	message := struct{ Error string }{Error: "Internal Server Error at " + request.URL.String() + ", using request method " + request.Method}
	jsonMessage, _ := json.Marshal(message)
	writer.WriteHeader(http.StatusInternalServerError)
	writer.Header().Add("Content-Type", "application/json")
	writer.Write(jsonMessage)
}

// parse path data into key/value pairings
func (base *Base) Parse(path string) map[string]interface{} {
	pairs := make(map[string]interface{}, 0)
	values := strings.Split(path, "/")
	for i, k := range values {
		if i < len(values) {
			pairs[k] = values[i+1]
		}
	}
	return pairs
}
