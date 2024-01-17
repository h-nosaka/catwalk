package fixtures

import (
	"github.com/google/uuid"
	"github.com/h-nosaka/catwalk/examples/models"
	"gorm.io/gorm"
)

func AccountPin(setters ...func(model *models.AccountPin)) *models.AccountPin {
	model := &models.AccountPin{
		Id: uuid.NewString()}
	for _, setter := range setters {
		setter(model)
	}
	return model
}

func CreateAccountPin(db *gorm.DB, setters ...func(model *models.AccountPin)) *models.AccountPin {
	model := AccountPin(setters...)
	if err := db.Create(model).Error; err != nil {
		return nil
	}
	return model
}
