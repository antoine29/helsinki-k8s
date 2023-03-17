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

	if natsSubject, exists := os.LookupEnv("NATS_SUBJECT"); !exists {
		log.Println("NATS_SUBJECT env var not set")
		return
	} else {
	  log.Printf("Publishing to: %s nats subject\n", natsSubject)
  }

	server.Run(port)
}
