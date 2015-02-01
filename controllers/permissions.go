package controllers

import (
	"net/http"
)

type Permissions struct {
	Base
}

// registration
func (base *Base) RegisterWithRouter(
	addRoute func(
		uri string,
		callback func(
			writer http.ResponseWriter,
			request *http.Request),
		methods ...string)) {
	addRoute("/permissions", base.CreateGroup, "PUT")
	addRoute("/permissions", base.CreateGroup, "POST")
	addRoute("/permissions", base.CreateGroup, "DELETE")
	addRoute("/permissions/action", base.AddAction, "POST")
	addRoute("/permissions/action", base.DeleteAction, "DELETE")

}

func (permissions *Permissions) CreateGroup(writer http.ResponseWriter, request *http.Request) {
	// @todo(casey): create group
}

func (permissions *Permissions) EditGroup(writer http.ResponseWriter, request *http.Request) {
	// @todo(casey): update group
}

func (permissions *Permissions) DeleteGroup(writer http.ResponseWriter, request *http.Request) {
	// @todo(casey): delete group (disables, does not actually delete)
}

func (permissions *Permissions) AddAction(writer http.ResponseWriter, request *http.Request) {
	// @todo(casey): add action from group
}

func (permissions *Permissions) DeleteAction(writer http.ResponseWriter, request *http.Request) {
	// @todo(casey): delete action from group
}

/**
 *
 * The list of available actions may later be moved into another part of the system
 * doing so would allow us to store karma costs and other meta-data such as the
 * related subsystem and even dependencies, but that depends on a 100% dynamic build
 *
 */
