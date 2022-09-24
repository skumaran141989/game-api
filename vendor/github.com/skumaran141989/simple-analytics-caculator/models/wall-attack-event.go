package models

type WallAttackEvent struct {
	Event
	wallId   string
	strength int64
}

func (e *WallAttackEvent) GetWallId() string {
	return e.wallId
}

func (e *WallAttackEvent) SetWallId(wallId string) {
	e.wallId = wallId
}

func (e *WallAttackEvent) GetStrength() int64 {
	return e.strength
}

func (e *WallAttackEvent) SetStrength(strength int64) {
	e.strength = strength
}
