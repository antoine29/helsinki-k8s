package main

import (
	"log"
	"os"

	"github.com/antoine29/todos-queue-consumer/pkg/natsClient"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(err.Error())
	}

	var (
		natsSubject            string
		isNatsSubjectEnvVarSet bool
		natsUrl                string
		isNatsUrlEnvVarSet     bool
	)

	if natsSubject, isNatsSubjectEnvVarSet = os.LookupEnv("NATS_SUBJECT"); !isNatsSubjectEnvVarSet {
		log.Println("'NATS_SUBJECT' env var not set. Exiting.")
		return
	}

	if natsUrl, isNatsUrlEnvVarSet = os.LookupEnv("NATS_URL"); !isNatsUrlEnvVarSet {
		log.Println("'NATS_URL' env var not set. Exiting.")
		return
	}

	natsClient.Subscribe(natsUrl, natsSubject)
}
