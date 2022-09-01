package src

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func DownloadFile(URL, fileName string) error {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("received non 200 response code")
	}
	//Create a empty file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func LocalFileExists(name string) bool {
	files, err := ioutil.ReadDir("/tmp/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.Name() == name && !file.IsDir() {
			return true
		}
	}

	return false
}

func GetFile(name string) string {
	if LocalFileExists(name) {
		return ReadFile(name)
	}

	DownloadFile("url", name)
	return ReadFile(name)
}

func ReadFile(name string) string {
	return name
}
