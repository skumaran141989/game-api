package models

type Base struct {
	CreatedDate int64  `json:"created_date"`
	CreatedBy   string `json:"created_by"`
	UpdatedDate int64  `json:"updated_date"`
	UpdatedBy   string `json:"updated_by"`
	DeletedDate int64  `json:"deleted_date"`
	DeletedBy   string `json:"deleted_by"`
}

func (e *Base) GetCreatedDate() int64 {
	return e.CreatedDate
}

func (e *Base) SetCreatedDate(CreatedDate int64) {
	e.CreatedDate = CreatedDate
}

func (e *Base) GetCreatedBy() string {
	return e.CreatedBy
}

func (e *Base) SetCreatedBy(CreatedBy string) {
	e.CreatedBy = CreatedBy
}

func (e *Base) GetUpdatedDate() int64 {
	return e.CreatedDate
}

func (e *Base) SetUpdatedDate(UpdatedDate int64) {
	e.UpdatedDate = UpdatedDate
}

func (e *Base) GetUpdatedBy() string {
	return e.UpdatedBy
}

func (e *Base) SetUpdatedBy(UpdatedBy string) {
	e.UpdatedBy = UpdatedBy
}

func (e *Base) GetDeletedDate() int64 {
	return e.CreatedDate
}

func (e *Base) SetDeletedDate(DeletedDate int64) {
	e.DeletedDate = DeletedDate
}

func (e *Base) GetDeletedBy() string {
	return e.CreatedBy
}

func (e *Base) SetDeletedBy(DeletedBy string) {
	e.DeletedBy = DeletedBy
}
