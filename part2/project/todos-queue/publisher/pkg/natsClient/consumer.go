package natsClient

import (
	"log"

	nats "github.com/nats-io/nats.go"
)

func Subscribe() {
	nconn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Printf("Error connecting to nats\n%s\n", err.Error())
		return
	}

	nchan := make(chan *nats.Msg, 64)
	nsub, _ := nconn.ChanSubscribe("foo", nchan)

	for nmsg := range nchan {
		log.Println(string(nmsg.Data))
	}

	// Unsubscribe
	nsub.Unsubscribe()
}
