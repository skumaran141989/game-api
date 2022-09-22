package models

type Event struct {
	id     string `gorm:"primary_key;default:uuid_v4()"`
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
