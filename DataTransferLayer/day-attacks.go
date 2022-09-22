package DataTransferLayer

type DayAttacks struct {
	Attacks []*Attack `json:"attacks"`
}

func (e *DayAttacks) GetAttacks() []*Attack {
	return e.Attacks
}

func (e *DayAttacks) SetAttacks(dayAttack []*Attack) {
	e.Attacks = dayAttack
}
