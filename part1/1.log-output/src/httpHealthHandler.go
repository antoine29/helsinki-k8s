package src

import (
	"fmt"
	"log"
	"net/http"
)

func GetHttpStatus(url string) bool {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println("error hiting:", url)
		log.Fatalln(err)
		return false
	}

	return true
}
