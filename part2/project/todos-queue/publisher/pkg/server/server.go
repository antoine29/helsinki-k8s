package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/antoine29/todos-queue-publisher/pkg/models"
	"github.com/antoine29/todos-queue-publisher/pkg/natsClient"
)

var natsSubject string = os.Getenv("SUBJECT")

func publishController(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		WriteJsonResponse(res, http.StatusMethodNotAllowed, nil)
		return
	}

	payload := req.Body
	defer req.Body.Close()

	var message models.Message
	if err := json.NewDecoder(payload).Decode(&message); err != nil {
		WriteJsonResponse(res, http.StatusInternalServerError, BuildErrorResponse(err))
		return
	}

	var (
		jmessage []byte
		err      error
	)

	if jmessage, err = json.Marshal(message); err != nil {
		WriteJsonResponse(res, http.StatusInternalServerError, BuildErrorResponse(err))
		return
	}

	if err = natsClient.Publish(natsSubject, jmessage); err != nil {
		WriteJsonResponse(res, http.StatusInternalServerError, BuildErrorResponse(err))
		return
	}

	WriteJsonResponse(res, http.StatusInternalServerError, jmessage)
	return
}

func Run(port string) {
	log.Printf("Listening on: %s \n", port)
	http.HandleFunc("/api/publish", publishController)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
