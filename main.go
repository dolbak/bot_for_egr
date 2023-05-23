package main

import (
	"log"
	event_consumer "myGoApp/consumer/event-consumer"

	tgClient "myGoApp/clients/telegram"
	"myGoApp/events/telegram"
	"myGoApp/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, ""),
		files.New(storagePath),
	)

	log.Print("service started")

	conumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := conumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

/*func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}*/
