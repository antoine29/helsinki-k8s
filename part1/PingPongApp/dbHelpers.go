package main

import (
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

var ExpectedDBenvVars = []string{
	"PG_HOST",
	"PG_PORT",
	"PG_USER",
	"PG_PASSWORD",
	"PG_DBNAME",
	"PG_SCHEMA",
	"PG_TABLE",
}

type Count struct {
	id         int
	count      int
	count_date string
	hash       string
}

func GotDBerror(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
		return true
	}

	return false
}

func FullDBEnvVars() bool {
	for _, dbEnvVarName := range ExpectedDBenvVars {
		value := os.Getenv(dbEnvVarName)
		if value == "" {
			fmt.Printf("Missing '%s' db env var", dbEnvVarName)
			return false
		}

		if dbEnvVarName == "PG_PORT" {
			if _, error := strconv.Atoi(os.Getenv("PG_PORT")); error != nil {
				fmt.Println("Error parsing PG_PORT env var")
				return false
			}
		}
	}

	return true
}
