package telegram

import (
	"myGoApp/clients/telegram"
	"myGoApp/events"
	"myGoApp/storage"
)

func (p *TelProcessor) chooseDay(chatID int) error {
	botState := events.AskWorkout
	storage.UpdateUserState(chatID, botState)

	var mon = telegram.NewInlineKeyboardButtonData("Пн", events.Monday)
	var tue = telegram.NewInlineKeyboardButtonData("Вт", events.Tuesday)
	var wen = telegram.NewInlineKeyboardButtonData("Ср", events.Wednesday)
	var firstRow = telegram.NewInlineKeyboardRow(mon, tue, wen)

	var thu = telegram.NewInlineKeyboardButtonData("Чт", events.Thursday)
	var fri = telegram.NewInlineKeyboardButtonData("Пт", events.Friday)
	var sat = telegram.NewInlineKeyboardButtonData("Сб", events.Saturday)
	var secondRow = telegram.NewInlineKeyboardRow(thu, fri, sat)

	var Sun = telegram.NewInlineKeyboardButtonData("Вс", events.Sunday)
	var thirdRow = telegram.NewInlineKeyboardRow(Sun)

	var week = telegram.NewInlineKeyboardMarkup(firstRow, secondRow, thirdRow)

	return p.tg.SendMessageWithKeyBoard(chatID, "Выберите день тренировки", week)
}

func (p *TelProcessor) chooseWorkout(chatID int, day string) error {
	botState := events.WorkoutIsPlaned
	storage.UpdateUserState(chatID, botState)
	//Тут должно быть сохранение(кэширование) дня недели
	//Тут должно быть создание нумированного списка тренировок
	var addNew = telegram.NewInlineKeyboardButtonData("Создать новую", "/workout")
	var row = telegram.NewInlineKeyboardRow(addNew)
	var keyboard = telegram.NewInlineKeyboardMarkup(row)
	return p.tg.SendMessageWithKeyBoard(chatID, "Вы выбрали: "+day+"\nВыберите тренировку или создайте новую.\n "+
		"(полагаю, тут нужно предоставить нумерованный  список тренировок, чел отправляет цифру нужного)"+
		"\n1. Тренировка 1"+
		"\n2. Тренировка 2...", keyboard)
}

func (p *TelProcessor) savePlan(chatID int, workout string) error {
	//Тут должно быть сохранение планирования
	return p.tg.SendMessage(chatID, "Вы запланировали тренировку")
}

// Начало создания новой тренировки
func (p *TelProcessor) chooseExercise(chatID int) error {
	botState := events.AskNumberOfSets
	storage.UpdateUserState(chatID, botState)
	//Должно быть создание нумированного списка упражнений
	var addNew = telegram.NewInlineKeyboardButtonData("Новое", "/exercise")
	var recommendation = telegram.NewInlineKeyboardButtonData("Рекомендации", "/recommendation")
	var row = telegram.NewInlineKeyboardRow(addNew, recommendation)
	var keyboard = telegram.NewInlineKeyboardMarkup(row)
	return p.tg.SendMessageWithKeyBoard(chatID, "\nВыберите упражнение или создайте новое.\n (полагаю, тут нужно предоставить список упражнений, чел отправляет цифру нужного)"+
		"\n1.Упражнение 1 \n2.Упражнение 2", keyboard)
}

// После выбора упражнения, создания нового или получения рекомендации добавляем информацию
func (p *TelProcessor) addSets(chatID int, exerciseName string) error {
	botState := events.AskNumberOfRepetitions
	storage.UpdateUserState(chatID, botState)
	//Тут должно быть сохранение(кэширование) названия
	return p.tg.SendMessage(chatID, "Вы выбрали упражнение "+exerciseName+". Сколько подходов?")
}

func (p *TelProcessor) addRepetitions(chatID int, numberOfSets string) error {
	botState := events.AskWeightInExercise
	storage.UpdateUserState(chatID, botState)
	//Тут должно быть сохранение(кэширование) количесва подходов
	return p.tg.SendMessage(chatID, "Сколько повторов?")
}

func (p *TelProcessor) addWeightInExercise(chatID int, repetitions string) error {
	botState := events.AskDescriptionOfWorkout
	storage.UpdateUserState(chatID, botState)
	//Тут должно быть сохранение(кэширование) количества пвторов
	return p.tg.SendMessage(chatID, "Сколько кг?")
}

// Сохранение упражнения, возможность добавить еще одно или написать комментарий к тренировке
func (p *TelProcessor) saveExerciseInfo(chatID int) error {
	botState := events.WorkoutIsCreated
	storage.UpdateUserState(chatID, botState)
	//Тут должно быть сохранение инфы по упражнению в тренировку

	var addNew = telegram.NewInlineKeyboardButtonData("Добавить еще", "/workout")
	var row = telegram.NewInlineKeyboardRow(addNew)
	var keyboard = telegram.NewInlineKeyboardMarkup(row)
	return p.tg.SendMessageWithKeyBoard(chatID, "Вы добавили упражнение в тренировку. Напишите комментарий к тренировке", keyboard)
}

func (p *TelProcessor) saveNewWorkout(chatID int) error {
	botState := events.WaitingForSomething
	storage.UpdateUserState(chatID, botState)
	//Тут должно быть сохранение тренировки. Если она создавалась при планировании, добавить ее на нужный день.
	return p.tg.SendMessage(chatID, "Вы создали тренировку.")
}

func (p *TelProcessor) func1(chatID int) error {
	return p.tg.SendMessage(chatID, "Вы создали упражнение")
}

/*func (p *TelProcessor) addTime(chatID int) error {
	return p.tg.SendMessage(chatID, "Сколько времени на выполнение упражнения?")
}*/

func (p *TelProcessor) createExercise(chatID int) error {
	return p.tg.SendMessage(chatID, "Вы создали упражнение")
}

func (p *TelProcessor) recommendations(chatId int) error {
	p.tg.SendMessage(chatId, "Выберете день тренировки")
	p.tg.SendMessage(chatId, "У вас уже есть тренировка на этот день. Хотите изменить ее?")
	p.tg.SendMessage(chatId, "Хотите получить рекомендацию?")
	return p.tg.SendMessage(chatId, "Готово")
}
