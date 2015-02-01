package main

import (
	"net/http"

	"./controllers"

	"github.com/cdelorme/go-log"
	// "github.com/cdelorme/go-config"
	// "github.com/cdelorme/go-option"
	"github.com/cdelorme/go-routing"
)

func main() {

	// prepare logger
	logger := log.Logger{Level: log.Info}
	logger.Info("Emptiness is everything.")

	// load configuration
	// conf, confErr := config.Load("config.json")
	// logger.Debug("error loading config: %s", confErr)

	// @todo(casey): prepare cli options
	// @todo(casey): load options and merge with config
	// @todo(casey): acquire default settings

	// temporarily assume 127.0.0.1:8080 for testing
	Address := "127.0.0.1:8080"
	// @note: port 80 is reserved on osx, and won't be allowed

	// prepare router
	router := routing.Router{}

	// @todo(casey):prepare controlers
	permissionsController := controllers.Permissions{}

	// @todo(casey):register controllers with router
	router.RegisterController(&permissionsController)

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
