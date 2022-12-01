package main

import (
	"fmt"
	"net/http"
	"os"
)

func statusController(w http.ResponseWriter, req *http.Request) {
	runMode := getRunMode()
	status, counter := GetCurrentStatus(runMode)
	WriteStatus(status, counter, runMode)
	fmt.Println(req.RequestURI)
	fmt.Fprintln(w, status)
}

func healthController(w http.ResponseWriter, req *http.Request) {
	dbReadiness := DBreadiness()
	if dbReadiness {
		fmt.Fprintln(w, "DB is ready")
		return
	}

	http.Error(w, "Error connecting to DB", 503)
}

func Server(port string) {
	http.HandleFunc("/", statusController)
	http.HandleFunc("/health", healthController)

	fmt.Printf("Listening on: %s \n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func getRunMode() string {
	mode := os.Getenv("GO_RUNMODE")

	if mode == "" {
		fmt.Println("Running in memory mode")
		return "memory"
	}

	if mode != "db" {
		fmt.Printf("Invalid '%s' mode. Running in memory mode \n", mode)
		return "memory"
	}

	fmt.Println("Running in db mode")
	return "db"
}
