package events

type Fatcher interface {
	Fetch(limit int) ([]Event, error)
}

type Processor interface {
	Process(e Event) error
}

type Type int

const (
	Unknown Type = iota
	Message
	CallbackQuery
)

type Event struct {
	Type Type
	Text string
	Meta interface{}
}

type State int

const (
	WaitingForSomething State = iota
	StartIvent
	HelpEvent
	AskDay
	AskWorkout
	WorkoutIsPlaned
	AskExercise
	AskNumberOfSets
	AskNumberOfRepetitions
	AskWeightInExercise
	AskTime
	AskAnotherExercise
	AskCommentToWorkout
	WorkoutIsCreated
	AskNameOfNewExercise
	AskDescriptionOfExercise
	ExerciseIdCreated
	AskCommentToExercise
	AskRecommendation
	AskTypeOfWorkout
	AskComplexityOfExercise
	SendRecommendation
)

type Weekdays string

const (
	Monday    = "Понедельник"
	Tuesday   = "Вторник"
	Wednesday = "Среда"
	Thursday  = "Четверг"
	Friday    = "Пятница"
	Saturday  = "Суббота"
	Sunday    = "Воскресенье"
)
