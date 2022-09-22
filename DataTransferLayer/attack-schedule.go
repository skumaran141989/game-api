package DataTransferLayer

type Schedule struct {
	DayAttacks []*DayAttacks `json:"day_attacks"`
}

func (e *Schedule) GetDayAttacks() []*DayAttacks {
	return e.DayAttacks
}

func (e *Schedule) SetDayAttacks(dayAttack []*DayAttacks) {
	e.DayAttacks = dayAttack
}

func (e *Schedule) GetDayAttack(index int) *DayAttacks {
	return e.DayAttacks[index]
}
