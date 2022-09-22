package models

type User struct {
	Id        string `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Base
}

func (e *User) GetId() string {
	return e.Id
}

func (e *User) SetId(id string) {
	e.Id = id
}

func (e *User) GetFirstName() string {
	return e.FirstName
}

func (e *User) SetFirstName(firstName string) {
	e.FirstName = firstName
}

func (e *User) GetLastName() string {
	return e.LastName
}

func (e *User) SetLastName(lastName string) {
	e.LastName = lastName
}
