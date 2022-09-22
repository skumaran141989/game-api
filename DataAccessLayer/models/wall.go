package models

type Wall struct {
	Id        string `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Direction string `json:"direction"`
	Height    int64  `json:"height"`
	Base
}

func (e *Wall) GetId() string {
	return e.Id
}

func (e *Wall) SetId(id string) {
	e.Id = id
}

func (e *Wall) GetDirection() string {
	return e.Direction
}

func (e *Wall) SetDirection(direction string) {
	e.Direction = direction
}

func (e *Wall) GetHeight() int64 {
	return e.Height
}

func (e *Wall) SetHeight(height int64) {
	e.Height = height
}
