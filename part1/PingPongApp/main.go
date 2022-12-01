package main

import (
	"fmt"
	"os"
)

func main() {
	runMode := getRunMode()
	isDBEnvVarsComplete := FullDBEnvVars()
	if runMode == "db" && !isDBEnvVarsComplete {
		fmt.Println("Setting runMode to 'memory'")
	}

	port := os.Getenv("GO_PORT")
	if port == "" {
		fmt.Println("Warning: 'GO_PORT' environment variable was not set, using 8080 as default.")
		port = "8080"
	}

	Server(port)
}
