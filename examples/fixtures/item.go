package fixtures

import (
	"github.com/google/uuid"
	"github.com/h-nosaka/catwalk/examples/models"
	"gorm.io/gorm"
)

func Item(setters ...func(model *models.Item)) *models.Item {
	model := &models.Item{
		Id: uuid.NewString()}
	for _, setter := range setters {
		setter(model)
	}
	return model
}

func CreateItem(db *gorm.DB, setters ...func(model *models.Item)) *models.Item {
	model := Item(setters...)
	if err := db.Create(model).Error; err != nil {
		return nil
	}
	return model
}
