package natsClient

import (
	"log"
	"os"

	tg "github.com/antoine29/todos-queue-telegram-client/pkg"
	nats "github.com/nats-io/nats.go"
  _ "github.com/joho/godotenv/autoload"
)

var TG_BOT_TOKEN, isTgBotTokenEnvVarSet = os.LookupEnv("TG_BOT_TOKEN")
var TG_CHANNEL_NAME, isTgChannelNameEnvVarSet = os.LookupEnv("TG_CHANNEL_NAME")
var NATS_URL, isNatsUrlEnvVarSet = os.LookupEnv("NATS_URL")

func Subscribe(subject string) {
  if !isNatsUrlEnvVarSet {
    log.Println("'NATS_URL' env var not set. Exiting.")
    return
  }

	nconn, err := nats.Connect(nats.DefaultURL)
	defer nconn.Close()
	if err != nil {
		log.Printf("Error connecting to nats\n%s\n", err.Error())
		return
	}

	natsChannel := make(chan *nats.Msg, 64)
	nsub, err := nconn.ChanSubscribe(subject, natsChannel)
	defer nsub.Unsubscribe()
	if err != nil {
		log.Printf("Error subscribing to nats subject. \n%s\n", err.Error())
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

