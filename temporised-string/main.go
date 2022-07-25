package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	//fmt.Println("hii")
	interval, _ := strconv.ParseInt(os.Args[1], 0, 64)
	//fmt.Println("integer interval: ", interval)
	/*
		params := os.Args[1:]
		for i, param := range params {
			fmt.Println(i, param)
			fmt.Println("type: ", fmt.Sprintf("%T", param))
		}
	*/

	intervalString := fmt.Sprintf("%ds", interval)
	//fmt.Println("interval string: ", intervalString)

	duration, _ := time.ParseDuration(intervalString)
	ticker := time.NewTicker(duration)
	for tick := range ticker.C {
		fmt.Println("Tick: ", tick)
	}
}
