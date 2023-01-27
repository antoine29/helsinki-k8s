package src

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var EnvVarsDict map[string]string = make(map[string]string)

func LoadEnvVarsDict(silentMode bool) {
	envVarsDict, err := godotenv.Read()
	if err == nil {
		if !silentMode {
			log.Print(".env file loaded")
		}

		EnvVarsDict = envVarsDict
	} else {
		EnvVarsDict["PG_HOST"] = os.Getenv("PG_HOST")
		EnvVarsDict["PG_PORT"] = os.Getenv("PG_PORT")
		EnvVarsDict["PG_USER"] = os.Getenv("PG_USER")
		EnvVarsDict["PG_PASSWORD"] = os.Getenv("PG_PASSWORD")
		EnvVarsDict["PG_DBNAME"] = os.Getenv("PG_DBNAME")
		EnvVarsDict["PG_SCHEMA"] = os.Getenv("PG_SCHEMA")
		EnvVarsDict["GO_PORT"] = os.Getenv("GO_PORT")
	}

	alternativeEnvVarsDict, err := godotenv.Read("/config/.env")
	if err == nil {
		if !silentMode {
			log.Print("alternative /config/.env file loaded. Overwriting PG_PASSWORD value")
		}

		EnvVarsDict["PG_PASSWORD"] = alternativeEnvVarsDict["PG_PASSWORD"]
	}
}
