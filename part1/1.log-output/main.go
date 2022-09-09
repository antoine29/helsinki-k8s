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

		fmt.Println(status)
		statusMemoHandler.SetStatus(status)
		statusFileHandler.SetStatus(status)
	}
}

func main() {
	helpers.TryToReadEnvFiles()
	paramsDict := helpers.BuildProgramParamsDict(os.Args[1:])

	if helpers.IsParamPassed("reader", paramsDict) && !helpers.IsParamPassed("writer", paramsDict) {
		helpers.CheckMandatoryParams(paramsDict, helpers.ReaderExpectedParams)
		serverPort := paramsDict["serverPort"]
		fmt.Println("Running in reader mode")
		server.Server(serverPort, paramsDict)
		return
	}

	if !helpers.IsParamPassed("reader", paramsDict) && helpers.IsParamPassed("writer", paramsDict) {
		helpers.CheckMandatoryParams(paramsDict, helpers.WriterExpectedParams)
		secsInterval := helpers.ParseStrToInt(paramsDict["secsInterval"])
		randomStrLength := helpers.ParseStrToInt(paramsDict["strLength"])
		fmt.Println("Running in writer mode")
		generateRandomStrings(secsInterval, randomStrLength)
		return
	}

	helpers.CheckMandatoryParams(paramsDict, append(helpers.WriterExpectedParams, helpers.ReaderExpectedParams...))
	serverPort := paramsDict["serverPort"]
	secsInterval := helpers.ParseStrToInt(paramsDict["secsInterval"])
	randomStrLength := helpers.ParseStrToInt(paramsDict["strLength"])
	fmt.Println("Running in writer/reader mode")
	go server.Server(serverPort, paramsDict)
	generateRandomStrings(secsInterval, randomStrLength)
}
