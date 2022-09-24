package models

type Event struct {
	id     string
	userId string
	Base
}

func (e *Event) GetId() string {
	return e.id
}

func (e *Event) SetId(id string) {
	e.id = id
}

func (e *Event) GetUserId() string {
	return e.userId
}

func (e *Event) SetUserId(userId string) {
	e.userId = userId
}
