package natsClient

import (
	nats "github.com/nats-io/nats.go"
)

func Publish(url string, subject string, message []byte) error {
	nconn, err := nats.Connect(url)
	defer nconn.Close()

	if err != nil {
		return err
	}

	if err := nconn.Publish(subject, message); err != nil {
		return err
	}

	return nil
}
