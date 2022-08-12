package main

import (
	"fmt"
	"net/http"
)

var counter int = 0

func current(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, counter)
	counter++
}

func Server(port string) {
	http.HandleFunc("/", current)

	fmt.Printf("Listening on: %s \n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
