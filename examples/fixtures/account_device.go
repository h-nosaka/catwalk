package fixtures

import (
	"github.com/google/uuid"
	"github.com/h-nosaka/catwalk/examples/models"
	"gorm.io/gorm"
)

func AccountDevice(setters ...func(model *models.AccountDevice)) *models.AccountDevice {
	model := &models.AccountDevice{
		Id: uuid.NewString()}
	for _, setter := range setters {
		setter(model)
	}
	return model
}

func CreateAccountDevice(db *gorm.DB, setters ...func(model *models.AccountDevice)) *models.AccountDevice {
	model := AccountDevice(setters...)
	if err := db.Create(model).Error; err != nil {
		return nil
	}
	return model
}
