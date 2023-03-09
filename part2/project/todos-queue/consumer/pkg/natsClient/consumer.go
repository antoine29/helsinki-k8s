package natsClient

import (
	"log"

	nats "github.com/nats-io/nats.go"
)

func Subscribe(subject string) {
	nconn, err := nats.Connect(nats.DefaultURL)
	defer nconn.Close()
	if err != nil {
		log.Printf("Error connecting to nats\n%s\n", err.Error())
		return
	}

	nchan := make(chan *nats.Msg, 64)
	nsub, err := nconn.ChanSubscribe(subject, nchan)
	defer nsub.Unsubscribe()
	if err != nil {
		log.Printf("Error connecting to subscribing\n%s\n", err.Error())
		return
	}

	log.Println("Subscribed to", subject)
	for nmsg := range nchan {
		log.Println(string(nmsg.Data))
	}
}
