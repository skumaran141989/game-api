package models

type Tribe struct {
	User
	TribeType string `json:"tribe_type"`
}

func (e *Tribe) GetTribeType() string {
	return e.TribeType
}

func (e *Tribe) SetTribeType(tribeType string) {
	e.TribeType = tribeType
}
