package pkg

import (
	"fmt"
	"log"
  "net/url"
)

func SendTgMessageFromBotToChannel(tgBotToken string, tgChannelName string, message string) error {
  encodedMsg := url.QueryEscape(message)
  url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s", tgBotToken, tgChannelName, encodedMsg)
  if err := HttpGet(url); err != nil {
    log.Println("Error sending TG message", "\n", err.Error())
  } else {
    log.Println("TG message sent:", message)
  }

  return nil
}

