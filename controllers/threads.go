package controllers

import (
	"net/http"
)

type Threads struct {
	Base
}

// registration
func (threads *Threads) RegisterWithRouter(
	addRoute func(
		uri string,
		callback func(
			writer http.ResponseWriter,
			request *http.Request),
		methods ...string)) {
	addRoute("/threads", threads.CreateThread, "PUT")
	addRoute("/thread", threads.EditThread, "POST")
	addRoute("/thread", threads.DeleteThread, "DELETE")
	addRoute("/thread", threads.GetThread, "GET")
	addRoute("/threads", threads.GetThreads, "GET")
	addRoute("/thread/posts", threads.GetPostsByThread, "GET")
	addRoute("/threads/recent", threads.GetRecentThreads, "GET")
	addRoute("/threads/active", threads.GetActiveThreads, "GET")
	addRoute("/threads/member", threads.GetPostsByMember, "GET")
	addRoute("/posts", threads.CreatePost, "PUT")
	addRoute("/post", threads.EditPost, "POST")
	addRoute("/post", threads.DeletePost, "DELETE")
	addRoute("/posts/member", threads.GetThreadsByMember, "GET")
}

func (threads *Threads) CreateThread(writer http.ResponseWriter, request *http.Request) {
	// @todo(casey): create a thread
}

func (threads *Threads) EditThread(writer http.ResponseWriter, request *http.Request) {
	// @todo(casey): edit a thread
}

func (threads *Threads) DeleteThread(writer http.ResponseWriter, request *http.Request) {
	// @todo(casey): delete a thread
}

func (threads *Threads) GetThread(writer http.ResponseWriter, request *http.Request) {
	// @todo(casey): get a thread
}

func (threads *Threads) GetThreads(writer http.ResponseWriter, request *http.Request) {
	// @todo(casey): get threads
}

func (threads *Threads) GetRecentThreads(writer http.ResponseWriter, request *http.Request) {
	// @todo(casey): get most recent threads (by `created`)
}

func (threads *Threads) GetActiveThreads(writer http.ResponseWriter, request *http.Request) {
	// @todo(casey): get active threads (by `updated`)
}

func (threads *Threads) CreatePost(writer http.ResponseWriter, request *http.Request) {
	// @todo(casey): create a post
}

func (threads *Threads) EditPost(writer http.ResponseWriter, request *http.Request) {
	// @todo(casey): edit a post
}

func (threads *Threads) DeletePost(writer http.ResponseWriter, request *http.Request) {
	// @todo(casey): delete a post
}

func (threads *Threads) GetPostsByThread(writer http.ResponseWriter, request *http.Request) {
	// @todo(casey): get posts by thread id
}

func (threads *Threads) GetPostsByMember(writer http.ResponseWriter, request *http.Request) {
	// @todo(casey): get posts by member id
}

func (threads *Threads) GetThreadsByMember(writer http.ResponseWriter, request *http.Request) {
	// @todo(casey): get threads by member id
}
