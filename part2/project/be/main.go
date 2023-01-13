package main

import (
	"antoine29/go/web-server/src/router"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	TryToReadEnvFiles()
	port := os.Getenv("GO_PORT")
	if port == "" {
		fmt.Println("Warning: 'GO_PORT' environment variable was not set, using 8080 as default.")
		port = "8080"
	}

	server := router.SetupServer()

	fmt.Printf("Go to: 'http://localhost:%s/swagger/index.html' to check Swagger API docs.\n", port)
	server.Run(fmt.Sprintf(":%s", port))
}

func TryToReadEnvFiles() {
	readingLocalEnvError := godotenv.Load()
	if readingLocalEnvError == nil {
		fmt.Println("Reading .env file")
		return
	}

	fmt.Println("Couldm't find .env file")
}
