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

	if subject, exists := os.LookupEnv("SUBJECT"); !exists {
		log.Println("SUBJECT env var not set")
		return
	} else {
		natsClient.Subscribe(subject)
	}
}
