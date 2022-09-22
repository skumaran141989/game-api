package repos

import (
	models "github.com/game-api/DataAccessLayer/models"
	"github.com/jinzhu/gorm"
)

type WallAttackEvents struct {
	db *gorm.DB
}

type WallAttackEventDB interface {
	Get(id string) (*models.WallAttackEvent, error)
	GetAll(query string) ([]models.WallAttackEvent, error)
	Update(t models.WallAttackEvent)
	Delete(id string)
}

func NewWallAttackEvent(db *gorm.DB) *WallAttackEvents {
	return &WallAttackEvents{
		db: db,
	}
}

func (e *WallAttackEvents) Get(id string) (*models.WallAttackEvent, error) {
	var event models.WallAttackEvent
	e.db.Table("WallAttackEvents").Where("id = " + id).First(&event)

	return &event, nil
}

func (e *WallAttackEvents) GetAll(query string) ([]models.WallAttackEvent, error) {
	var events []models.WallAttackEvent
	err := e.db.Table("WallAttackEvents").Where(query).Order("created_at").Find(&events).Error
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (e *WallAttackEvents) Update(event models.WallAttackEvent) error {
	err := e.db.Table("WallAttackEvents").Save(&event).Error
	if err != nil {
		return err
	}

	return nil
}

func (e *WallAttackEvents) Delete(id string) error {
	err := e.db.Table("WallAttackEvents").Delete("id = " + id).Error
	if err != nil {
		return err
	}

	return nil
}
