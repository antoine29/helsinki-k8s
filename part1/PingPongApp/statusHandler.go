package main

import (
	"fmt"
	"time"
)

func GetCurrentStatus(runMode string) (string, int) {
	var counter int

	if runMode == "memory" {
		counter = ReadFromFile()
	}

	if runMode == "db" {
		count := *ReadCountFromDB()
		counter = count.count
	}

	timeStamp := time.Now()
	status := fmt.Sprintf("time: %s\nPing / Pongs: %d\n", timeStamp, counter)
	return status, counter
}

func WriteStatus(status string, counter int, runMode string) {
	if runMode == "memory" {
		WriteToFile(status)
	}

	if runMode == "db" {
		WriteToDB(counter, status)
	}
}
