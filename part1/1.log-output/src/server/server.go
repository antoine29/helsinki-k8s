package server

import (
	helpers "antoine29/go/log-output/src"

	statusFileHandler "antoine29/go/log-output/src/statusHandlers/inFile"
	statusHttpHandler "antoine29/go/log-output/src/statusHandlers/inHttpReq"
	statusMemoHandler "antoine29/go/log-output/src/statusHandlers/inMemo"
	"fmt"
	"net/http"
)

// refactor this controller handlers to avoid repeating them (generics?)
func inMemoStatusEndpoint(w http.ResponseWriter, req *http.Request) {
	currentStatus := statusMemoHandler.GetStatus()
	fmt.Printf("/status/memory:\n%s\n", currentStatus)
	fmt.Fprintln(w, currentStatus)
}

func inFileStatusEndpoint(w http.ResponseWriter, req *http.Request) {
	currentStatus := statusFileHandler.GetStatus()
	fmt.Printf("/status/file:\n%s\n", currentStatus)
	fmt.Fprintln(w, currentStatus)
}

func okStatusEndpoint(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
    http.NotFound(w, req)
  }

	w.WriteHeader(200)
  return 
}

func Server(port string, programParams map[string]string) {
	http.HandleFunc("/", okStatusEndpoint)
	http.HandleFunc("/status/memory", inMemoStatusEndpoint)
	http.HandleFunc("/status/file", inFileStatusEndpoint)
	if helpers.IsParamPassed("url", programParams) {
		url := programParams["url"]
		fmt.Printf("Using '%s' as targuet url for '/status/http' endpoint\n", url)
		http.HandleFunc("/status/http", func(w http.ResponseWriter, req *http.Request) {
			httpResponse := statusHttpHandler.GetStatus(url)
			currentStatus := httpResponse

			envMessage, error := helpers.GetMessageEnvVar()
			if !error {
				currentStatus = envMessage + "\n" + currentStatus
			}

			fmt.Printf("/status/http:\n%s\n", currentStatus)
			fmt.Fprintln(w, currentStatus)
		})
	}

	fmt.Printf("Listening on: %s \n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
