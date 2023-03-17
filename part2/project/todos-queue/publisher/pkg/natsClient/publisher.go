package natsClient

import (
	"errors"
	"os"

	_ "github.com/joho/godotenv/autoload"
	nats "github.com/nats-io/nats.go"
)

var NATS_URL, isNatsUrlEnvVarSet = os.LookupEnv("NATS_URL")

func Publish(subject string, message []byte) error {
  if !isNatsUrlEnvVarSet {
    errors.New("'NATS_URL' env var not set. Exiting.")
  }

	nconn, err := nats.Connect(nats.DefaultURL)
	defer nconn.Close()

	if err != nil {
		return err
	}

	if err := nconn.Publish(subject, message); err != nil {
		return err
	}

	return nil
}

