package storage

import "myGoApp/events"

type Storage interface {
	Save(p *Messege) error
	IsExist(p *Messege) (bool, error)
}
type Messege struct {
	Text     string
	UserName string
}

func UpdateUserState(ID int, state events.State) {
	UserStateMap[ID] = state
}

var UserStateMap = make(map[int]events.State)

type Messege1 struct {
	Text     string
	Func func()
}
