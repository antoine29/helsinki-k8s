package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

var dbEnvVarsMap = map[string]string{
	"host":     "PG_HOST",
	"port":     "PG_PORT",
	"user":     "PG_USER",
	"password": "PG_PASSWORD",
	"dbName":   "PG_DBNAME",
	"dbSchema": "PG_SCHEMA",
	"dbTable":  "PG_TABLE",
}

var (
	host     = os.Getenv(dbEnvVarsMap["host"])
	port, _  = strconv.Atoi(os.Getenv(dbEnvVarsMap["port"]))
	user     = os.Getenv(dbEnvVarsMap["user"])
	password = os.Getenv(dbEnvVarsMap["password"])
	dbname   = os.Getenv(dbEnvVarsMap["dbName"])
	schema   = os.Getenv(dbEnvVarsMap["dbSchema"])
	table    = os.Getenv(dbEnvVarsMap["dbTable"])
)

type Count struct {
	id         int
	count      int
	count_date string
	hash       string
}

func getDBconn() (*sql.DB, bool) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, dbOpenError := sql.Open("postgres", psqlconn)
	if gotDBerror(dbOpenError) {
		return nil, true
	}

	dbPingError := db.Ping()
	if gotDBerror(dbPingError) {
		return nil, true
	}

	return db, false
}

func getCounts() []Count {
	counts := make([]Count, 0)

	db, dbError := getDBconn()
	if dbError {
		return counts
	}

	rows, readError := db.Query(fmt.Sprintf("SELECT id, count, count_date, hash FROM %s.%s", schema, table))
	if gotDBerror(readError) {
		return counts
	}

	defer db.Close()
	defer rows.Close()

	for rows.Next() {
		_count := Count{}

		parsingRowError := rows.Scan(&_count.id, &_count.count, &_count.count_date, &_count.hash)
		if !gotDBerror(parsingRowError) {
			counts = append(counts, _count)
		}
	}

	return counts
	// fmt.Printf("%+v\n", counts)
}

func insertCount(count int, hash string) {
	db, dbError := getDBconn()
	if dbError {
		return
	}

	defer db.Close()
	currentTime := time.Now().Format("01-02-2006  15:04:05")
	insertQuery := fmt.Sprintf("INSERT INTO %s.%s (count, count_date, hash) values (%d, '%s', '%s')", schema, table, count, currentTime, hash)
	_, insertError := db.Exec(insertQuery)
	gotDBerror(insertError)
}

// todo: move to helpers
func gotDBerror(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
		return true
	}

	return false
}

func WriteToDB(count int, status string) {
	insertCount(count, status)
}

func FullDBEnvVars() bool {
	for _, dbEnvVarName := range dbEnvVarsMap {
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
