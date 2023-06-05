package telegram

import (
	"log"
	"myGoApp/events"
	"myGoApp/storage"
	"strings"
)

const (
	StartCmd          = "/start"
	HelpCmd           = "/help"
	PlanCmd           = "/plan"
	CreateWorkout     = "/workout"
	CreateExercise    = "/exercise"
	GetRecommendation = "/recommendation"
)

func (p *TelProcessor) processInputMessage(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s' ID: '%d", text, username, chatID)
	var botState events.State
	switch text {
	case StartCmd:
		botState = events.StartIvent
	case HelpCmd:
		botState = events.HelpEvent
	case PlanCmd:
		botState = events.AskDay
	case CreateWorkout:
		botState = events.AskExercise

	case CreateExercise:
		botState = events.AskNameOfNewExercise
	case GetRecommendation:
		botState = events.AskRecommendation
	//этого тут быть не должно
	case "AskExercise":
		botState = events.AskExercise
	default:
		botState = storage.UserStateMap[chatID]

	}
	storage.UpdateUserState(chatID, botState)
	p.processState(botState, text, chatID)
	return nil
}

func (p *TelProcessor) processState(state events.State, text string, ID int) error {
	switch state {
	case events.WaitingForSomething:
		return p.unidentifiedAction(ID)
	case events.StartIvent:
		return p.sendHello(ID)
	case events.HelpEvent:
		return p.sendHelp(ID)
	case events.AskDay:
		return p.chooseDay(ID)
	case events.AskWorkout:
		return p.chooseWorkout(ID, text)
	case events.WorkoutIsPlaned:
		return p.savePlan(ID, text)
	case events.AskExercise:
		return p.chooseExercise(ID)
	case events.AskNumberOfSets:
		return p.addSets(ID, text)
	case events.AskNumberOfRepetitions:
		return p.addRepetitions(ID, text)
	case events.AskWeightInExercise:
		return p.addWeightInExercise(ID, text)
	case events.AskDescriptionOfWorkout:
		return p.saveExerciseInfo(ID)
	case events.WorkoutIsCreated:
		return p.saveNewWorkout(ID)
	default:
		botState := events.WaitingForSomething
		storage.UpdateUserState(ID, botState)
		return p.tg.SendMessage(ID, "Состояние не сделано")
	}
}
