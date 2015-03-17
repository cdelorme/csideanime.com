package main

import (
	"net/http"

	"github.com/cdelorme/go-log"
	// "github.com/cdelorme/go-config"
	// "github.com/cdelorme/go-option"
	"github.com/cdelorme/go-routing"

	"./controllers"
)

func main() {

	// prepare logger
	logger := log.Logger{Level: log.Info}

	// load configuration
	// conf, confErr := config.Load("config.json")
	// logger.Debug("error loading config: %s", confErr)

	// @todo(casey): prepare cli options
	// @todo(casey): load options and merge with config
	// @todo(casey): acquire default settings

	// temporarily assume 127.0.0.1:9010 for testing
	Address := "127.0.0.1:9010"
	// @note: port 80 is reserved on osx and won't run

	// prepare router /w alpha api prefix
	router := routing.Router{Prefix: "/api/alpha"}

	// prepare controlers
	tagsController := controllers.Tags{}
	permissionsController := controllers.Permissions{}
	membersController := controllers.Members{}
	threadsController := controllers.Threads{}

	// register controllers with router
	router.RegisterController(&tagsController)
	router.RegisterController(&permissionsController)
	router.RegisterController(&membersController)
	router.RegisterController(&threadsController)

	// display router info
	logger.Info("Router: %+v", router)

	// load server
	server := http.Server{
		Addr:           Address,
		MaxHeaderBytes: 1 << 20,
		Handler:        &router,
	}

	// initialize server
	logger.Debug("Server halted: %s", server.ListenAndServe())
}
