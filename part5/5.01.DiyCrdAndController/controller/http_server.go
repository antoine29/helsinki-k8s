package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, %s!", r.URL.Path[1:])
}

func Listen() {
	http.HandleFunc("/api/deployments", handler)

	port := os.Getenv("GO_PORT")
	if port == "" {
		fmt.Println("Error: 'GO_PORT' environment variable not set, using 8080 as defult.")
		port = "8080"
	}

	fmt.Println("Server started in port:", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// this kinda watch all dummy sites
// but it does it by requesting k8s api directly
// can we use the k8s go client instead?
func ListenDummySites() {
	resp, err := http.Get("http://localhost:8080/apis/stable.anth/v1/dummysites?watch=true")
	if err != nil {
		log.Println("Error hiting  url")
	}

	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			log.Println("Error reading bytes")
			log.Println(err.Error())
		}

		log.Println(string(line))
	}
}
