package fixtures

import (
	"github.com/google/uuid"
	"github.com/h-nosaka/catwalk/examples/postgresql/models"
	"gorm.io/gorm"
)

func AccountActivate(setter func(model *models.AccountActivate)) *models.AccountActivate {
	model := &models.AccountActivate{
		Id: uuid.NewString(),
	}
	setter(model)
	return model
}

func CreateAccountActivate(db *gorm.DB, setter func(model *models.AccountActivate)) *models.AccountActivate {
	model := AccountActivate(setter)
	if err := db.Create(model).Error; err != nil {
		return nil
	}
	return model
}
