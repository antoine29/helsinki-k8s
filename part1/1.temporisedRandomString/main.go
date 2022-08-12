package main

import (
	"fmt"
	"os"
	"strconv"
)

func parseStrToInt(str string) int {
	intVal, parseError := strconv.ParseInt(str, 0, 0)
	if parseError != nil {
		fmt.Printf("Error: Cannot parse '%s' to int", str)
		os.Exit(3)
	}

	return int(intVal)
}

func main() {
	paramsDict := BuildProgramParamsDict(os.Args[1:])

	for _, expectedParam := range ExpectedParams {
		_, exists := paramsDict[expectedParam]
		if !exists {
			fmt.Printf("Error: Expected '%s' parameter. \n", expectedParam)
			os.Exit(3)
		}
	}

	secsInterval := parseStrToInt(paramsDict["secsInterval"])
	randomStrLength := parseStrToInt(paramsDict["strLength"])
	serverPort := paramsDict["serverPort"]

	go Server(serverPort)
	PrintRandomString(randomStrLength, secsInterval)
}
