package fixtures

import (
	"github.com/google/uuid"
	"github.com/h-nosaka/catwalk/examples/postgresql/models"
	"gorm.io/gorm"
)

func AccountPincode(setter func(model *models.AccountPincode)) *models.AccountPincode {
	model := &models.AccountPincode{
		Id: uuid.NewString(),
	}
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
