package repos

import (
	"github.com/game-api/DataAccessLayer/models"
	"github.com/jinzhu/gorm"
)

type Tribes struct {
	db *gorm.DB
}

type TribeDB interface {
	Get(id string) (*models.Tribe, error)
	GetAll(id string) ([]models.Tribe, error)
	Update(tribe models.Tribe)
	Delete(id string)
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
