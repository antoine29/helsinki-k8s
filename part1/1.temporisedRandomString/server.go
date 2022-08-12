package main

import (
	"fmt"
	"net/http"
)

func current(w http.ResponseWriter, req *http.Request) {
	currentStatus := GetStatus()
	response := fmt.Sprintf("%s", currentStatus)
	fmt.Fprintln(w, response)

}

func Server() {
	http.HandleFunc("/current", current)

	port := 8090
	fmt.Printf("Listening on: %d \n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
