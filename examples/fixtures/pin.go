package fixtures

import (
	"github.com/google/uuid"
	"github.com/h-nosaka/catwalk/examples/models"
	"gorm.io/gorm"
)

func Pin(setters ...func(model *models.Pin)) *models.Pin {
	model := &models.Pin{
		Id: uuid.NewString()}
	for _, setter := range setters {
		setter(model)
	}
	return model
}

func CreatePin(db *gorm.DB, setters ...func(model *models.Pin)) *models.Pin {
	model := Pin(setters...)
	if err := db.Create(model).Error; err != nil {
		return nil
	}
	return model
}
