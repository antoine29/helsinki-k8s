package main

import (
	"fmt"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	url := os.Getenv("URL")

	if port == "" {
		fmt.Println("Error: 'PORT' environment variable not set, using 8080 as defult.")
		port = "8080"
	}

	if url == "" {
		panic("Error: 'URL' environment variable not set.")
	}

	Listen(port, url)

	// PORT=8083 URL=https://example.com/ go run .
}
