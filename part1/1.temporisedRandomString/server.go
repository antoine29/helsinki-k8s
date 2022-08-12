package main

import (
	"fmt"
	"net/http"
)

func current(w http.ResponseWriter, req *http.Request) {
	currentStatus := GetStatus()
	fmt.Fprintln(w, currentStatus)
}

func Server(port string) {
	http.HandleFunc("/current", current)

	fmt.Printf("Listening on: %s \n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
