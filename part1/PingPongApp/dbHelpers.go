package main

import (
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

var DbEnvVarsMap = map[string]string{
	"host":     "PG_HOST",
	"port":     "PG_PORT",
	"user":     "PG_USER",
	"password": "PG_PASSWORD",
	"dbName":   "PG_DBNAME",
	"dbSchema": "PG_SCHEMA",
	"dbTable":  "PG_TABLE",
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
	for _, dbEnvVarName := range DbEnvVarsMap {
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
