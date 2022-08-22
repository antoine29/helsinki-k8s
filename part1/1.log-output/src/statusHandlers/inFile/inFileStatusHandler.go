package inFile

import (
	"crypto/md5"
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
	file, e0 := os.Create(filePath)
	handleError(e0)

	_, e1 := file.WriteString(_status)
	handleError(e1)

	file.Sync()
	file.Close()
}

func GetStatus() string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Sprintf("Error reading from file: %s \n%s", filePath, err)
	}

	hash := hashFile(data)
	return fmt.Sprintf("MD5 file hash: %s \n%s", hash, string(data))
}

func hashFile(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))
}
