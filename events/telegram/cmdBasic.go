package telegram

import (
	"myGoApp/events"
	"myGoApp/storage"
)

func (p *TelProcessor) sendHello(chatID int) error {
	botState := events.WaitingForSomething
	storage.UpdateUserState(chatID, botState)
	return p.tg.SendMessage(chatID, "Этот бот не умеет отправлять твое сообщение с твоим именем или сообщение об повторном соощении \nА еще тут есть команды: \n"+
		"/start \n"+
		"/help \n"+
		"/plan позваляет запланировать тренировку\n"+
		"/workout позволяет создать новую тренировку\n"+
		"/exercise позволяет создать новое упражнение\n"+
		"/recommendation получите новое упражнение")
}

func (p *TelProcessor) sendHelp(chatID int) error {
	botState := events.WaitingForSomething
	storage.UpdateUserState(chatID, botState)
	return p.tg.SendMessage(chatID, "Доступные в боте команды:\n"+
		"/start \n"+
		"/help \n"+
		"/plan позваляет запланировать тренировку\n"+
		"/workout позволяет создать новую тренировку\n"+
		"/exercise позволяет создать новое упражнение\n"+
		"/recommendation получите новое упражнение")
}

func (p *TelProcessor) unidentifiedAction(chatID int) error {
	return p.tg.SendMessage(chatID, "Неопознанное действие, введите команду из списка: \n/start \n/help \n/plan\n/workout\n/exercise\n/recommendation")
}
