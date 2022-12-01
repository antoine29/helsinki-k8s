package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

var (
	host     = os.Getenv(DbEnvVarsMap["host"])
	port, _  = strconv.Atoi(os.Getenv(DbEnvVarsMap["port"]))
	user     = os.Getenv(DbEnvVarsMap["user"])
	password = os.Getenv(DbEnvVarsMap["password"])
	dbname   = os.Getenv(DbEnvVarsMap["dbName"])
	schema   = os.Getenv(DbEnvVarsMap["dbSchema"])
	table    = os.Getenv(DbEnvVarsMap["dbTable"])
)

func getDBconn() (*sql.DB, bool) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, dbOpenError := sql.Open("postgres", psqlconn)
	if GotDBerror(dbOpenError) {
		return nil, true
	}

	dbPingError := db.Ping()
	if GotDBerror(dbPingError) {
		return nil, true
	}

	return db, false
}

func DBreadiness() bool {
	db, dbConnError := getDBconn()
	if dbConnError {
		return false
	}

	_, dbReadError := db.Query(fmt.Sprintf("SELECT COUNT(1) FROM %s.%s", schema, table))
	if GotDBerror(dbReadError) {
		return false
	}

	return true
}

func getCounts() []Count {
	counts := make([]Count, 0)

	db, dbError := getDBconn()
	if dbError {
		return counts
	}

	rows, readError := db.Query(fmt.Sprintf("SELECT id, count, count_date, hash FROM %s.%s", schema, table))
	if GotDBerror(readError) {
		return counts
	}

	defer db.Close()
	defer rows.Close()

	for rows.Next() {
		_count := Count{}

		parsingRowError := rows.Scan(&_count.id, &_count.count, &_count.count_date, &_count.hash)
		if !GotDBerror(parsingRowError) {
			counts = append(counts, _count)
		}
	}

	return counts
	// fmt.Printf("%+v\n", counts)
}

func ReadCountFromDB() *Count {
	db, dbError := getDBconn()
	if dbError {
		return nil
	}

	row, readError := db.Query(fmt.Sprintf("SELECT * FROM %s.%s WHERE count = (select MAX(count) FROM %s.%s)", schema, table, schema, table))
	if GotDBerror(readError) {
		return nil
	}

	defer db.Close()
	defer row.Close()

	count := Count{}
	for row.Next() {
		parsingRowError := row.Scan(&count.id, &count.count, &count.count_date, &count.hash)
		// fmt.Println(count)
		if GotDBerror(parsingRowError) {
			log.Println("Error parsing count row")
			return nil
		}
	}

	return &count
	// fmt.Printf("%+v\n", counts)
}

func WriteToDB(count int, hash string) {
	db, dbError := getDBconn()
	if dbError {
		return
	}

	defer db.Close()
	currentTime := time.Now().Format("01-02-2006  15:04:05")
	insertQuery := fmt.Sprintf("INSERT INTO %s.%s (count, count_date, hash) values (%d, '%s', '%s')", schema, table, count+1, currentTime, hash)
	_, insertError := db.Exec(insertQuery)
	GotDBerror(insertError)
}
