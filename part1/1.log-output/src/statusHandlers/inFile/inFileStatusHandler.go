package inFile

import (
	"fmt"
	"os"
)

var filePath = "/tmp/status"

func handleError(e error) {
	if e != nil {
		panic(e)
	}
}

func SetStatus(_status string) {
	f, e0 := os.Create(filePath)
	handleError(e0)

	_, e1 := f.WriteString(_status)
	handleError(e1)

	f.Sync()
	f.Close()
}

func GetStatus() string {
	dat, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Sprintf("Error reading from file: %s \n%s", filePath, err)
	}

	return string(dat)
}
