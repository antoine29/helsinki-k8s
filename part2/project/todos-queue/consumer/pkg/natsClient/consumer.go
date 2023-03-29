package natsClient

import (
	"log"
	"os"

	tg "github.com/antoine29/todos-queue-telegram-client/pkg"
	_ "github.com/joho/godotenv/autoload"
	nats "github.com/nats-io/nats.go"
)

var TG_BOT_TOKEN, isTgBotTokenEnvVarSet = os.LookupEnv("TG_BOT_TOKEN")
var TG_CHANNEL_NAME, isTgChannelNameEnvVarSet = os.LookupEnv("TG_CHANNEL_NAME")

func Subscribe(url string, subject string) {
	nconn, err := nats.Connect(url)
	defer nconn.Close()
	if err != nil {
		log.Printf("Error connecting to nats url: %s\n%s", url, err.Error())
		return
	}

	natsChannel := make(chan *nats.Msg, 64)
	nsub, err := nconn.QueueSubscribeSyncWithChan(subject, "queue"+subject, natsChannel)
	defer nsub.Unsubscribe()
	if err != nil {
		log.Printf("Error subscribing to nats subject: %s\n%s\n", subject, err.Error())
		return
	}

	log.Println("Subscribed to", subject)
	for natsMsg := range natsChannel {
		msg := string(natsMsg.Data)
		log.Println("Message received:", msg)
		sendMsgToTg(msg)
	}
}

func sendMsgToTg(message string) {
	if isTgBotTokenEnvVarSet && isTgChannelNameEnvVarSet {
		if err := tg.SendTgMessageFromBotToChannel(TG_BOT_TOKEN, TG_CHANNEL_NAME, message); err != nil {
			log.Println(err.Error())
		}
	} else {
		log.Println("TG env vars not set")
	}
}
