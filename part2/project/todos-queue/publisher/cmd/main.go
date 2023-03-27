package main

import (
	"log"
	"os"

	"github.com/antoine29/todos-queue-publisher/pkg/server"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(err.Error())
	}

	port := "8080"
	if envVarPort, exists := os.LookupEnv("PORT"); exists {
		port = envVarPort
	} else {
		log.Println("PORT env var not set")
	}

	log.Printf("Listening on: %s \n", port)

	var (
		natsUrl            string
		isNatsUrlEnvVarSet bool
		natsSubject        string
		exists             bool
	)

	if natsSubject, exists = os.LookupEnv("NATS_SUBJECT"); !exists {
		log.Println("NATS_SUBJECT env var not set")
		return
	}

	if natsUrl, isNatsUrlEnvVarSet = os.LookupEnv("NATS_URL"); !isNatsUrlEnvVarSet {
		log.Println("NATS_URL env var not set")
		return
	}

	log.Printf("Publishing to nats_url: %s \t nats_subject: %s \n", natsUrl, natsSubject)

	server.Run(port)
}
