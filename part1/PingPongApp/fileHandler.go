package main

import (
	"os"
)

// todo: get this path from env
var filePath = "/tmp/status"
var counter int = 0

func handleError(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFromFile() int {
	counter++
	return counter
}

func WriteToFile(content string) {
	file, fileCreationError := os.Create(filePath)
	handleError(fileCreationError)

	_, fileWrittingError := file.WriteString(content)
	handleError(fileWrittingError)

	file.Sync()
	file.Close()
}
