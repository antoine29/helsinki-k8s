package main

import (
	"antoine29/go/web-server/src/router"
	"fmt"
	"os"
)

func main() {
	port := os.Getenv("GO_PORT")
	if port == "" {
		fmt.Println("Warning: 'GO_PORT' environment variable was not set, using 8080 as default.")
		port = "8080"
	}

	r := router.SetupRouter()

	fmt.Printf("Go to: 'http://localhost:%s/swagger/index.html' to check Swagger API docs.\n", port)
	r.Run(fmt.Sprintf(":%s", port))
}
