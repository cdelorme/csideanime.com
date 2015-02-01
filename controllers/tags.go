package controllers

import (
	"net/http"
)

type Tags struct {
	Base
}

// registration
func (tags *Tags) RegisterWithRouter(
	addRoute func(
		uri string,
		callback func(
			writer http.ResponseWriter,
			request *http.Request),
		methods ...string)) {
	addRoute("/tags", tags.CreateTag, "PUT")
	addRoute("/tags", tags.EditTag, "POST")
	addRoute("/tags", tags.GetTags, "GET")
	addRoute("/tags", tags.DeleteTag, "DELETE")
}

func (tags *Tags) CreateTag(writer http.ResponseWriter, request *http.Request) {
	// @todo(casey): create new tag
}

func (tags *Tags) EditTag(writer http.ResponseWriter, request *http.Request) {
	// @todo(casey): edit a tag
}

func (tags *Tags) GetTags(writer http.ResponseWriter, request *http.Request) {
	// @todo(casey): delete a tag
}

func (tags *Tags) DeleteTag(writer http.ResponseWriter, request *http.Request) {
	// @todo(casey): delete a tag
}
