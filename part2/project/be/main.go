package main

import (
	config "antoine29/go/web-server/src"
	"antoine29/go/web-server/src/router"
	"fmt"
	"log"
)

func main() {
	config.LoadEnvVarsDict(false)
	port, portEnvVarIsSet := config.EnvVarsDict["GO_PORT"]
	if portEnvVarIsSet {
		log.Println("Warning: 'GO_PORT' environment variable was not set, using 8080 as default.")
		port = "8080"
	}

	server := router.SetupServer()

	log.Printf("Go to: 'http://localhost:%s/swagger/index.html' to check Swagger API docs.\n", port)
	server.Run(fmt.Sprintf(":%s", port))
}
