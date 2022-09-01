package inFile

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func SetStatus(_status string) {
	fmt.Println("'SetStatus' not defined for inHttpReq handler")
}

func GetStatus(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error hiting")
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	stringBody := string(body)
	hashedStringBody := hashFile([]byte(stringBody))
	return fmt.Sprintf("%s\n%s", hashedStringBody, stringBody)
}

func hashFile(data []byte) string {
	return fmt.Sprintf("%x", md5.Sum(data))
}
