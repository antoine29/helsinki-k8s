package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func generateRandomString() string {
	seed := time.Now().UnixNano()
	randomSource := rand.NewSource(seed)
	generator := rand.New(randomSource)
	randomLength := generator.Intn(10) + 1
	randomString := ""
	for i := 0; i < randomLength; i++ {
		// lowercase letters: 97 - 122 => 0+97 = 97, 25+97 = 122
		randomLowerCaseAsciInt := generator.Intn(26) + 97
		randomLowerCaseAsci := string(rune(randomLowerCaseAsciInt))
		// fmt.Println("randomLowerCaseAsci", randomLowerCaseAsci)
		randomString += randomLowerCaseAsci
	}

	return randomString
}

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

	stringInterval := fmt.Sprintf("%ds", intInterval)
	randomString := generateRandomString()

	fmt.Println("random string:", randomString, "interval to print: ", stringInterval)

	tickerDuration, _ := time.ParseDuration(stringInterval)
	ticker := time.NewTicker(tickerDuration)
	for tick := range ticker.C {
		fmt.Println("Tick: ", tick, randomString)
	}
}
