package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/antoine29/todos-queue-publisher/pkg/models"
	"github.com/antoine29/todos-queue-publisher/pkg/natsClient"
	_ "github.com/joho/godotenv/autoload"
)

var NATS_URL = os.Getenv("NATS_URL")
var NATS_SUBJECT = os.Getenv("NATS_SUBJECT")

func publishController(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		SendJsonResponse(res, http.StatusMethodNotAllowed, nil)
		return
	}

	payload := req.Body
	defer req.Body.Close()

	var message models.TodoMessage
	if err := json.NewDecoder(payload).Decode(&message); err != nil {
		SendJsonResponse(res, http.StatusInternalServerError, BuildErrorResponse(err))
		return
	}

	var (
		jmessage []byte
		err      error
	)

	if jmessage, err = json.Marshal(message); err != nil {
		SendJsonResponse(res, http.StatusInternalServerError, BuildErrorResponse(err))
		return
	}

	if err = natsClient.Publish(NATS_URL, NATS_SUBJECT, jmessage); err != nil {
		SendJsonResponse(res, http.StatusInternalServerError, BuildErrorResponse(err))
		return
	}

	SendJsonResponse(res, http.StatusOK, jmessage)
	log.Println("Sent: ", string(jmessage))
	return
}

func Run(port string) {
	http.HandleFunc("/api/publish", publishController)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}