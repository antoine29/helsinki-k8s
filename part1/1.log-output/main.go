package main

import (
	"fmt"
	"os"
	"time"

	helpers "antoine29/go/log-output/src"
	server "antoine29/go/log-output/src/server"
	statusFileHandler "antoine29/go/log-output/src/statusHandlers/inFile"
	statusMemoHandler "antoine29/go/log-output/src/statusHandlers/inMemo"
)

func generateRandomStrings(secsInterval int, randomStrLength int) {
	stringInterval := fmt.Sprintf("%ds", secsInterval)
	tickerDuration, _ := time.ParseDuration(stringInterval)
	ticker := time.NewTicker(tickerDuration)
	for tick := range ticker.C {
		randomString := helpers.GenerateRandomString(randomStrLength)
		status := fmt.Sprintf("%s %s", tick.String(), randomString)

		fmt.Println(tick, status)
		statusMemoHandler.SetStatus(status)
		statusFileHandler.SetStatus(status)
	}
}

func main() {
	paramsDict := helpers.BuildProgramParamsDict(os.Args[1:])
	helpers.CheckPassedParams(paramsDict)

	secsInterval := helpers.ParseStrToInt(paramsDict["secsInterval"])
	randomStrLength := helpers.ParseStrToInt(paramsDict["strLength"])
	serverPort := paramsDict["serverPort"]

	go server.Server(serverPort)
	generateRandomStrings(secsInterval, randomStrLength)
}
