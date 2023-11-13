package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Listen(port, url string) {
	http.HandleFunc("/", handlerBuilder(url))

	fmt.Println("Server started in port:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getHtml(url string) string {
	// todo: add validation for valid urls
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	content, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func handlerBuilder(url string) func(w http.ResponseWriter, r *http.Request) {
	html := getHtml(url)

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, html)
	}
}
