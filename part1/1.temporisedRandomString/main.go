package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	/*
		programParams := os.Args[1:]
		for i, param := range programParams {
			fmt.Println(i, param)
			fmt.Println("type: ", fmt.Sprintf("%T", param))
		}
	*/

	if len(os.Args) == 1 || len(os.Args) > 2 {
		fmt.Println("Error: You have to pass a ms integer")
		os.Exit(3)
	}

	intervalParam := os.Args[1]
	intInterval, parseError := strconv.ParseInt(intervalParam, 0, 64)
	if parseError != nil {
		fmt.Println("Error: Cannot parse", intervalParam, "to int")
		os.Exit(3)
	}

	go Server()
	PrintRandomString(7, intInterval)
}
