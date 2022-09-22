package DataTransferLayer

type Attack struct {
	TribeId  string `json:"tribe_id"`
	WallId   string `json:"wall_id"`
	Strength int64  `json:"strength"`
}

func (e *Attack) GetTribeId() string {
	return e.TribeId
}

func (e *Attack) SetTribeId(tribeId string) {
	e.TribeId = tribeId
}

func (e *Attack) GetWallId() string {
	return e.WallId
}

func (e *Attack) SetWallId(wallId string) {
	e.WallId = wallId
}

func (e *Attack) GetStrength() int64 {
	return e.Strength
}

func (e *Attack) SetStrength(strength int64) {
	e.Strength = strength
}
