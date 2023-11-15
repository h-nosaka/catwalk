package fixtures

import (
	"github.com/h-nosaka/catwalk/examples/mysql/models"
	"gorm.io/gorm"
)

func AccountPincode(setters ...func(model *models.AccountPincode)) *models.AccountPincode {
	model := &models.AccountPincode{}
	for _, setter := range setters {
		setter(model)
	}
	return model
}

func CreateAccountPincode(db *gorm.DB, setters ...func(model *models.AccountPincode)) *models.AccountPincode {
	model := AccountPincode(setters...)
	if err := db.Create(model).Error; err != nil {
		return nil
	}
	return model
}
