package controllers

import (
	"net/http"
)

type Members struct {
	Base
}

// registration
func (members *Members) RegisterWithRouter(
	addRoute func(
		uri string,
		callback func(
			writer http.ResponseWriter,
			request *http.Request),
		methods ...string)) {
	addRoute("/members", members.CreateMember, "PUT")
	addRoute("/members", members.EditMember, "POST")
	addRoute("/members", members.DeleteMember, "DELETE")
	addRoute("/members", members.GetMember, "GET")
	addRoute("/members/login", members.Login, "POST")
	addRoute("/members/profile", members.GetMemberProfile, "GET")
	addRoute("/members/profiles", members.GetMemberProfiles, "GET")
}

func (members *Members) CreateMember(writer http.ResponseWriter, request *http.Request) {
	// @todo(casey): create new member
}

func (members *Members) EditMember(writer http.ResponseWriter, request *http.Request) {
	// @todo(casey): edit member
}

func (members *Members) DeleteMember(writer http.ResponseWriter, request *http.Request) {
	// @todo(casey): delete/disable member
}

func (members *Members) GetMember(writer http.ResponseWriter, request *http.Request) {
	// @todo(casey): Get a single member
}

func (members *Members) GetMemberProfile(writer http.ResponseWriter, request *http.Request) {
	// @todo(casey): get a single members profile
}

func (members *Members) GetMemberProfiles(writer http.ResponseWriter, request *http.Request) {
	// @todo(casey): get all member profiles
}

func (members *Members) Login(writer http.ResponseWriter, request *http.Request) {
	// @todo(casey): login with username & password
}

// @note: this is very likely to be part of the base controller
func (members *Members) Checkin(writer http.ResponseWriter, request *http.Request) {
	// @todo(casey): checkin using auth components
}
