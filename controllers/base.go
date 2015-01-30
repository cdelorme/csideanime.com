package controllers

import (
	"net/http"
	"strings"
)

type Base struct {
	HandleBadRequest          func(writer http.ResponseWriter, request *http.Request)
	HandleInternalServerError func(writer http.ResponseWriter, request *http.Request)
}

// stub function (this one adds no routes)
func (base *Base) RegisterWithRouter(
	addRoute func(
		uri string,
		callback func(
			writer http.ResponseWriter,
			request *http.Request),
		methods ...string)) {
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
	for k, i := range values {
		if i < len(values) {
			pairs[k] = values[i+1]
		}
	}
	return pairs
}
