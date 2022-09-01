package main

import (
	"fmt"
	"net/http"
	"time"
)

var counter int = 0

func getCurrentStatus() string {
	counter++
	timeStamp := time.Now()
	status := fmt.Sprintf("time: %s\nPing / Pongs: %d\n", timeStamp, counter)
	return status
}

func statusController(w http.ResponseWriter, req *http.Request) {
	status := getCurrentStatus()
	WriteToFile(status)
	fmt.Println(req.RequestURI)
	fmt.Fprintln(w, status)
}

func Server(port string) {
	http.HandleFunc("/", statusController)

	fmt.Printf("Listening on: %s \n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
