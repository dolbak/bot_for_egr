package main

import (
	"log"
	event_consumer "myGoApp/consumer/event-consumer"
	"myGoApp/storage/files"

	tgClient "myGoApp/clients/telegram"
	"myGoApp/events/telegram"

)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, "5689795126:AAHBY1PfprRyNAWbVOyENpAydsi_x4awaMo"),
		files.New(storagePath),
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}
