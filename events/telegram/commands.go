package telegram

import (
	"log"
	"myGoApp/lib/e"
	"myGoApp/storage"
	"strings"
)

const (
	StartCmd = "/start"
	HelpCmd  = "/help"
)

func (p *TelProcessor) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s", text, username)

	switch text {
	case StartCmd:
		return p.sendHello(chatID)
	case HelpCmd:
		return p.tg.SendMessage(chatID, "Не жди помощи")
	default:
		return p.saveMess(chatID, text, username)
	}
}

func (p *TelProcessor) saveMess(chatID int, messegeText string, username string) (err error) {
	defer func() { err = e.WrapIfErr("can't do command: save message", err) }()

	messege := &storage.Messege{
		Text:     messegeText,
		UserName: username,
	}

	isExists, err := p.storage.IsExist(messege)
	if err != nil {
		return err
	}
	if isExists {
		return p.tg.SendMessage(chatID, "Вывести ты уже это писал")
	}

	if err := p.storage.Save(messege); err != nil {
		return err
	}

	if err := p.tg.SendMessage(chatID, messegeText+" ты "+username); err != nil {
		return err
	}

	return nil
}

func (p *TelProcessor) sendHello(chatID int) error {
	return p.tg.SendMessage(chatID, "Этот бот умеет отправлять твое сообщение с твоим именем или сообщение об повторном соощении \nа еще тут есть команда /help")
}
