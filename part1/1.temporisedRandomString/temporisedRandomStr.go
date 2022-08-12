package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomString(length int) string {
	seed := time.Now().UnixNano()
	randomSource := rand.NewSource(seed)
	generator := rand.New(randomSource)
	randomString := ""
	for i := 0; i < length; i++ {
		// lowercase letters: 97 - 122 => 0+97 = 97, 25+97 = 122
		randomLowerCaseAsciInt := generator.Intn(26) + 97
		randomLowerCaseAsci := string(rune(randomLowerCaseAsciInt))
		randomString += randomLowerCaseAsci
	}

	return randomString
}

func PrintRandomString(length int, msInterval int64) {
	stringInterval := fmt.Sprintf("%ds", msInterval)
	tickerDuration, _ := time.ParseDuration(stringInterval)
	ticker := time.NewTicker(tickerDuration)
	for tick := range ticker.C {
		randomString := generateRandomString(length)
		status := fmt.Sprintf("%s %s", tick.String(), randomString)
		SetStatus(status)
		fmt.Println(tick, status)
	}
}
