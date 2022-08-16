package server

import (
	statusFileHandler "antoine29/go/log-output/src/statusHandlers/inFile"
	statusMemoHandler "antoine29/go/log-output/src/statusHandlers/inMemo"
	"fmt"
	"net/http"
)

func inMemoStatusEndpoint(w http.ResponseWriter, req *http.Request) {
	currentStatus := statusMemoHandler.GetStatus()
	fmt.Fprintln(w, currentStatus)
}

func inFileStatusEndpoint(w http.ResponseWriter, req *http.Request) {
	currentStatus := statusFileHandler.GetStatus()
	fmt.Fprintln(w, currentStatus)
}

func Server(port string) {
	http.HandleFunc("/status/memory", inMemoStatusEndpoint)
	http.HandleFunc("/status/file", inFileStatusEndpoint)

	fmt.Printf("Listening on: %s \n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
