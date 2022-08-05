package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"antoine29/go/web-server/src/models"
	"antoine29/go/web-server/src/dao"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, %s!", r.URL.Path[1:])
}

func main() {
	// todo := models.ToDo{"0", "todo0", false}
	// fmt.Println(todo.Title)
	todo0 := src.AddToDo(models.ToDo{ Title: "ToDo0" })
	fmt.Println(todo0)
	todos := src.GetToDo_s()
	fmt.Println(todos)

	http.HandleFunc("/api/", handler)
	
	port := os.Getenv("GO_PORT")
	if port == "" {
		fmt.Println("Error: 'GO_PORT' environment variable not set, using 8080 as default.")
		port = "8080"
	}
	
	fmt.Println("Server started in port:", port)
	
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
