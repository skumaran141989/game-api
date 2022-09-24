package repos

import (
	"github.com/jinzhu/gorm"
	"github.com/skumaran141989/game-api/DataAccessLayer/models"
)

type Walls struct {
	db *gorm.DB
}

type WallDB interface {
	DB[models.Wall]
}

func NewWall(db *gorm.DB) *Walls {
	return &Walls{
		db: db,
	}
}

func (e *Walls) Get(id string) (*models.Wall, error) {
	var wall models.Wall
	e.db.Table("Walls").Where("id = " + id).First(&wall)

	return &wall, nil
}

func (e *Walls) GetAll(query string) ([]models.Wall, error) {
	var walls []models.Wall
	err := e.db.Table("Walls").Where(query).Find(&walls).Error
	if err != nil {
		return nil, err
	}

	return walls, nil
}

func (e *Walls) Update(wall *models.Wall) error {
	err := e.db.Table("Walls").Save(&wall).Error
	if err != nil {
		return err
	}

	return nil
}

func (e *Walls) Delete(id string) error {
	err := e.db.Table("Walls").Delete("id = " + id).Error
	if err != nil {
		return err
	}

	return nil
}
