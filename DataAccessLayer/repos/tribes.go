package repos

import (
	"github.com/jinzhu/gorm"
	"github.com/skumaran141989/game-api/DataAccessLayer/models"
)

type Tribes struct {
	db *gorm.DB
}

type TribeDB interface {
	DB[models.Tribe]
}

func NewTribe(db *gorm.DB) *Tribes {
	return &Tribes{
		db: db,
	}
}

func (e *Tribes) Get(id string) (*models.Tribe, error) {
	var tribe models.Tribe
	e.db.Table("Tribes").Where("id = " + id).First(&tribe)

	return &tribe, nil
}

func (e *Tribes) GetAll(query string) ([]models.Tribe, error) {
	var tribes []models.Tribe
	err := e.db.Table("Tribes").Where(query).Find(&tribes).Error
	if err != nil {
		return nil, err
	}

	return tribes, nil
}

func (e *Tribes) Update(tribe models.Tribe) error {
	err := e.db.Table("Tribes").Save(&tribe).Error
	if err != nil {
		return err
	}

	return nil
}

func (e *Tribes) Delete(id string) error {
	err := e.db.Table("Tribes").Delete("id = " + id).Error
	if err != nil {
		return err
	}

	return nil
}
