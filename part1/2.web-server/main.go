package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	
	port := os.Getenv("GO_PORT")
	if port == "" {
		fmt.Println("Error: 'GO_PORT' environment variable not set, using 8080 as defult.")
		port = "8080"
	}
	
	fmt.Println("Server started in port:", port)
	
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
