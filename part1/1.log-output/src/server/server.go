package server

import (
	helpers "antoine29/go/log-output/src"

	src "antoine29/go/log-output/src"
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

func Server(port string, programParams map[string]string) {
	http.HandleFunc("/status/memory", inMemoStatusEndpoint)
	http.HandleFunc("/status/file", inFileStatusEndpoint)
	http.HandleFunc("/status/http", func(w http.ResponseWriter, req *http.Request) {
		if helpers.IsParamPassed("url", programParams) {
			url := programParams["url"] + "/status"
			fmt.Printf("Using '%s' as targuet url\n", url)
			httpStatus := statusHttpHandler.GetStatus(url)
			envMessage, error := helpers.GetMessageEnvVar()
			if !error {
				httpStatus = envMessage + "\n" + httpStatus
			}

			fmt.Printf("/status/http:\n%s\n", httpStatus)
			fmt.Fprintln(w, httpStatus)
			return
		}

		http.Error(w, "/status/http endpoint is not set", 501)
	})
	http.HandleFunc("/http_health", func(w http.ResponseWriter, r *http.Request) {
		if helpers.IsParamPassed("url", programParams) {
			url := programParams["url"] + "/health"
			fmt.Printf("Using '%s' as targuet url\n", url)
			isHealthy := src.GetHttpStatus(url)
			if !isHealthy {
				http.Error(w, "Not healthy", 501)
				return
			}

			fmt.Fprintln(w, "healthy")
			return
		}

		http.Error(w, "/status/http endpoint is not set", 501)
	})

	fmt.Printf("Listening on: %s \n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
