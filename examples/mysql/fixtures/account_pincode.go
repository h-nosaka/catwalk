package fixtures

import (
	"github.com/h-nosaka/catwalk/examples/mysql/models"
	"gorm.io/gorm"
)

func AccountPincode(setter func(model *models.AccountPincode)) *models.AccountPincode {
	model := &models.AccountPincode{}
	setter(model)
	return model
}

func CreateAccountPincode(db *gorm.DB, setter func(model *models.AccountPincode)) *models.AccountPincode {
	model := AccountPincode(setter)
	if err := db.Create(model).Error; err != nil {
		return nil
	}
	return model
}
