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

func Server(port string, runMode string) {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		status := getCurrentStatus()
		if runMode == "memory" {
			WriteToFile(status)
		}

		if runMode == "db" {
			WriteToDB(counter, status)
		}

		fmt.Println(req.RequestURI)
		fmt.Fprintln(w, status)
	})

	fmt.Printf("Listening on: %s \n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
