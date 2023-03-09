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
	}

	if _, exists := os.LookupEnv("SUBJECT"); !exists {
		log.Println("SUBJECT env var not set")
		return
	}

	server.Run(port)
}
