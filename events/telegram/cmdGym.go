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
	return p.tg.SendMessageWithKeyBoard(chatID, "Вы выбрали: "+day+"\nВведите цифру нужной вам тренировки или создайте новую.\n "+
		"\n1. Тренировка 1"+
		"\n2. Тренировка 2...", keyboard)
}

func (p *TelProcessor) savePlan(chatID int, workout string) error {
	botState := events.WaitingForSomething
	storage.UpdateUserState(chatID, botState)
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
	return p.tg.SendMessageWithKeyBoard(chatID, "\nВведите цифру нужного вам упражнения или создайте новое.\n"+
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
	botState := events.AskTime
	storage.UpdateUserState(chatID, botState)
	//Тут должно быть сохранение(кэширование) количества пвторов
	return p.tg.SendMessage(chatID, "Сколько кг?")
}

func (p *TelProcessor) addTime(chatID int, weight string) error {
	botState := events.AskCommentToExercise
	storage.UpdateUserState(chatID, botState)
	//Тут должно быть кэширование веса
	return p.tg.SendMessage(chatID, "Введите время выполнения упражнения")
}
func (p *TelProcessor) addCommentToExercise(chatID int, time string) error {
	botState := events.AskAnotherExercise
	storage.UpdateUserState(chatID, botState)
	//Тут должно быть кэширование веса
	return p.tg.SendMessage(chatID, "Введите комментарий к упражнению")
}

// Сохранение упражнения
func (p *TelProcessor) saveExerciseInfo(chatID int, comment string) error {
	botState := events.AskCommentToWorkout
	storage.UpdateUserState(chatID, botState)
	//Тут должно быть сохранение инфы по упражнению в тренировку

	var addNew = telegram.NewInlineKeyboardButtonData("Добавить еще", "/workout")
	var not = telegram.NewInlineKeyboardButtonData("Нет", "_")
	var row = telegram.NewInlineKeyboardRow(addNew, not)
	var keyboard = telegram.NewInlineKeyboardMarkup(row)
	return p.tg.SendMessageWithKeyBoard(chatID, "Хотите добавить еще упражнение?", keyboard)
}

func (p *TelProcessor) addCommentToWorkout(chatID int) error {
	botState := events.WorkoutIsCreated
	storage.UpdateUserState(chatID, botState)
	return p.tg.SendMessage(chatID, "Введите комментарий к тренировке")
}

func (p *TelProcessor) saveNewWorkout(chatID int) error {
	botState := events.WaitingForSomething
	storage.UpdateUserState(chatID, botState)
	//Тут должно быть сохранение тренировки. Если она создавалась при планировании, добавить ее на нужный день.
	return p.tg.SendMessage(chatID, "Вы создали тренировку.")
}

func (p *TelProcessor) createNewExercise(chatID int) error {
	botState := events.AskDescriptionOfExercise
	storage.UpdateUserState(chatID, botState)
	return p.tg.SendMessage(chatID, "Введите название упражнения")
}

func (p *TelProcessor) addDescriptionOfExercise(chatID int, name string) error {
	botState := events.ExerciseIdCreated
	storage.UpdateUserState(chatID, botState)
	//должно быть кэширование названия
	return p.tg.SendMessage(chatID, "Введите описание упражнения")
}

func (p *TelProcessor) SaveNewExercise(chatID int, description string) error {
	botState := events.AskNumberOfSets
	storage.UpdateUserState(chatID, botState)
	var addInfo = telegram.NewInlineKeyboardButtonData("Добавить информацию", "_")
	var row = telegram.NewInlineKeyboardRow(addInfo)
	var keyboard = telegram.NewInlineKeyboardMarkup(row)
	//Сохранение упражнения, получение его ID
	return p.tg.SendMessageWithKeyBoard(chatID, "Вы создали упражнение", keyboard)
}

func (p *TelProcessor) chooseMuscleGroup(chatID int) error {
	botState := events.AskTypeOfWorkout
	storage.UpdateUserState(chatID, botState)

	var button1 = telegram.NewInlineKeyboardButtonData("Кнопки для выбора группы мышц", "_")
	var row = telegram.NewInlineKeyboardRow(button1)
	var keyboard = telegram.NewInlineKeyboardMarkup(row)

	return p.tg.SendMessageWithKeyBoard(chatID, "Выберите группу мышц", keyboard)
}

func (p *TelProcessor) chooseTypeOfWorkout(chatID int, muscleGroup string) error {
	botState := events.AskComplexityOfExercise
	storage.UpdateUserState(chatID, botState)
	//Тут должна кэшироваться группа мышц
	var button1 = telegram.NewInlineKeyboardButtonData("Кнопки для выбора выбора типа тренировки", "_")
	var row = telegram.NewInlineKeyboardRow(button1)
	var keyboard = telegram.NewInlineKeyboardMarkup(row)

	return p.tg.SendMessageWithKeyBoard(chatID, "Выберите тип тренировки", keyboard)
}

func (p *TelProcessor) chooseComplexity(chatID int, typeOfWorkout string) error {
	botState := events.SendRecommendation
	storage.UpdateUserState(chatID, botState)

	//Тут должен кэшироваться тип тренировки
	var button1 = telegram.NewInlineKeyboardButtonData("Кнопки для выбора сложности", "_")
	var row = telegram.NewInlineKeyboardRow(button1)
	var keyboard = telegram.NewInlineKeyboardMarkup(row)

	return p.tg.SendMessageWithKeyBoard(chatID, "Выберите сложность", keyboard)
}

func (p *TelProcessor) getRecommendation(chatID int, complexity string) error {
	botState := events.AskNumberOfSets
	storage.UpdateUserState(chatID, botState)

	//Тут должно выдаваться упражнение по предыдущим выборам
	var button1 = telegram.NewInlineKeyboardButtonData("Хочу", "_")
	var button2 = telegram.NewInlineKeyboardButtonData("Хочу получить другое", "/workout")
	var row = telegram.NewInlineKeyboardRow(button1, button2)
	var keyboard = telegram.NewInlineKeyboardMarkup(row)

	return p.tg.SendMessageWithKeyBoard(chatID, "Ваше упражнение: ...\nХотите его добавить в тренировку?", keyboard)
}

func (p *TelProcessor) func1(chatID int) error {
	return p.tg.SendMessage(chatID, "Вы создали упражнение")
}
