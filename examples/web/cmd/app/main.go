// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This program provides a sample web service that implements a
// RESTFul CRUD API against a MongoDB database.
package main

import (
	"os"
	"time"

	"github.com/blreynolds4/go-common/cfg"
	"github.com/blreynolds4/go-common/examples/web/cmd/app/routes"
	"github.com/blreynolds4/go-common/examples/web/internal/sys/app"
	"github.com/blreynolds4/go-common/log"
	"github.com/blreynolds4/go-common/web"
)

// init is called before main. We are using init to customize logging output.
func init() {

	// This is being added to showcase configuration.
	os.Setenv("KIT_LOGGING_LEVEL", "1")
}

// main is the entry point for the application.
func main() {

	// Initialize the configuration and logging systems. Plus anything
	// else the web app layer needs.
	app.Init("main", cfg.EnvProvider{Namespace: app.Namespace})

	// Check the environment for a configured port value.
	host := os.Getenv("HOST")
	if host == "" {
		host = ":3000"
	}

	log.User("main", "main", "Started : Listening on: %s", host)
	err := web.Run(host, routes.API(), time.Second, time.Second)
	log.User("main", "main", "Down : %v", err)
}
