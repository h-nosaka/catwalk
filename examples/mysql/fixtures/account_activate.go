package fixtures

import (
	"github.com/h-nosaka/catwalk/examples/mysql/models"
	"gorm.io/gorm"
)

func AccountActivate(setter func(model *models.AccountActivate)) *models.AccountActivate {
	model := &models.AccountActivate{}
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
