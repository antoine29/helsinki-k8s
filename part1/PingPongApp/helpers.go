package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func TryToReadEnvFiles() {
	readingLocalEnvError := godotenv.Load()
	if readingLocalEnvError == nil {
		fmt.Println("Reading .env file")
		return
	}

	fmt.Println("Couldn't find .env file")
}
